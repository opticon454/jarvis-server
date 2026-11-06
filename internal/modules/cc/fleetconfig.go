package cc

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"regexp"

	"github.com/alexberardi/jarvis-server/internal/platform/authn"
	"github.com/alexberardi/jarvis-server/internal/platform/httpx"
)

// Fleet config (docs/cc/14-fleet-config-and-recovery.md). The app seals a config once with the
// household key; jarvisd keeps the ciphertext and every node of the household pulls it, so a
// secret is entered once and a node added later receives it on its first pull. jarvisd never
// holds the key: the nudge carries no ciphertext and the list/pull routes return it sealed.

var fleetTypeRE = regexp.MustCompile(`^[a-z0-9][a-z0-9_.:-]{0,63}$`)

func fleetType(w http.ResponseWriter, r *http.Request) (string, bool) {
	t := r.PathValue("config_type")
	if !fleetTypeRE.MatchString(t) {
		validationError(w, "path -> config_type: String should match pattern '^[a-z0-9][a-z0-9_.:-]{0,63}$'")
		return "", false
	}
	return t, true
}

func (m *Module) handlePutFleetConfig(w http.ResponseWriter, r *http.Request, u authn.User) {
	hh := r.PathValue("household_id")
	ctype, ok := fleetType(w, r)
	if !ok {
		return
	}
	b, _, ok := readBody(w, r, false)
	if !ok {
		return
	}
	ct, _ := b.str("ciphertext", true)
	nonce, _ := b.str("nonce", true)
	tag, _ := b.str("tag", true)
	if !b.done(w) {
		return
	}
	ctx := r.Context()
	if err := m.requireRole(ctx, u.ID, hh, authn.RoleAdmin); err != nil {
		m.writeErr(w, err)
		return
	}
	var rev int64
	if err := m.deps.DB.Tx(ctx, func(tx *sql.Tx) error {
		return tx.QueryRowContext(ctx, `INSERT INTO cc_fleet_config (household_id, config_type, ciphertext, nonce, tag, revision, updated_at)
			VALUES (?, ?, ?, ?, ?, 1, ?)
			ON CONFLICT (household_id, config_type) DO UPDATE SET ciphertext = excluded.ciphertext, nonce = excluded.nonce,
				tag = excluded.tag, revision = cc_fleet_config.revision + 1, updated_at = excluded.updated_at
			RETURNING revision`, hh, ctype, ct, nonce, tag, dbTime(m.now())).Scan(&rev)
	}); err != nil {
		m.internalError(w, err)
		return
	}
	nodes, err := m.householdNodes(ctx, hh)
	if err != nil {
		m.internalError(w, err)
		return
	}
	for _, n := range nodes {
		// Nudge only; a node that misses it pulls on its next MQTT (re)connect.
		if err := m.bus.Publish(n.nodeID, "fleet/config", map[string]any{"household_id": hh, "config_type": ctype, "revision": rev}); err != nil {
			m.deps.Log.Warn("cc: fleet config nudge not published; node pulls on reconnect", "node", n.nodeID, "err", err)
		}
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"config_type": ctype, "revision": rev, "nodes_notified": len(nodes)})
}

func (m *Module) handleDeleteFleetConfig(w http.ResponseWriter, r *http.Request, u authn.User) {
	hh := r.PathValue("household_id")
	ctype, ok := fleetType(w, r)
	if !ok {
		return
	}
	ctx := r.Context()
	if err := m.requireRole(ctx, u.ID, hh, authn.RoleAdmin); err != nil {
		m.writeErr(w, err)
		return
	}
	// Acks go with the row so a later write restarting at revision 1 is not read as already applied.
	if err := m.deps.DB.Tx(ctx, func(tx *sql.Tx) error {
		if _, err := tx.ExecContext(ctx, `DELETE FROM cc_fleet_config_acks WHERE household_id = ? AND config_type = ?`, hh, ctype); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `DELETE FROM cc_fleet_config WHERE household_id = ? AND config_type = ?`, hh, ctype)
		return err
	}); err != nil {
		m.internalError(w, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"status": "deleted"})
}

func (m *Module) handleListFleetConfig(w http.ResponseWriter, r *http.Request, u authn.User) {
	hh := r.PathValue("household_id")
	ctx := r.Context()
	if err := m.requireRole(ctx, u.ID, hh, authn.RoleMember); err != nil {
		m.writeErr(w, err)
		return
	}
	out, err := m.fleetItems(ctx, `SELECT config_type, ciphertext, nonce, tag, revision FROM cc_fleet_config
		WHERE household_id = ? ORDER BY config_type`, hh, hh)
	if err != nil {
		m.internalError(w, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, out)
}

// handlePendingFleetConfig returns what this node has not applied: nodes are bound to their
// own path and household, like config push (D6).
func (m *Module) handlePendingFleetConfig(w http.ResponseWriter, r *http.Request, n *nodeCtx) {
	if n.ID != r.PathValue("node_id") {
		detail(w, http.StatusForbidden, "Cannot access other node's config")
		return
	}
	if n.HouseholdID == "" {
		httpx.WriteJSON(w, http.StatusOK, []map[string]any{})
		return
	}
	out, err := m.fleetItems(r.Context(), `SELECT f.config_type, f.ciphertext, f.nonce, f.tag, f.revision FROM cc_fleet_config f
		LEFT JOIN cc_fleet_config_acks a ON a.node_id = ? AND a.household_id = f.household_id AND a.config_type = f.config_type
		WHERE f.household_id = ? AND (a.revision IS NULL OR a.revision < f.revision) ORDER BY f.updated_at, f.config_type`, n.HouseholdID, n.ID, n.HouseholdID)
	if err != nil {
		m.internalError(w, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, out)
}

// fleetItems runs a query whose first bind arguments are listed in args and whose columns are
// config_type, ciphertext, nonce, tag, revision; hh is stamped on each item as household_id.
func (m *Module) fleetItems(ctx context.Context, query, hh string, args ...any) ([]map[string]any, error) {
	rows, err := m.deps.DB.Read.QueryContext(ctx, query, append([]any{}, args...)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []map[string]any{}
	for rows.Next() {
		var ctype, ct, nonce, tag string
		var rev int64
		if err := rows.Scan(&ctype, &ct, &nonce, &tag, &rev); err != nil {
			return nil, err
		}
		out = append(out, map[string]any{"household_id": hh, "config_type": ctype, "ciphertext": ct, "nonce": nonce, "tag": tag, "revision": rev})
	}
	return out, rows.Err()
}

func (m *Module) handleAckFleetConfig(w http.ResponseWriter, r *http.Request, n *nodeCtx) {
	if n.ID != r.PathValue("node_id") {
		detail(w, http.StatusForbidden, "Cannot access other node's config")
		return
	}
	ctype, ok := fleetType(w, r)
	if !ok {
		return
	}
	b, _, ok := readBody(w, r, false)
	if !ok {
		return
	}
	rev, _ := b.integer("revision", true)
	if !b.done(w) {
		return
	}
	ctx := r.Context()
	var cur int64
	err := m.deps.DB.Read.QueryRowContext(ctx, `SELECT revision FROM cc_fleet_config WHERE household_id = ? AND config_type = ?`,
		n.HouseholdID, ctype).Scan(&cur)
	if errors.Is(err, sql.ErrNoRows) {
		detail(w, http.StatusNotFound, "Fleet config not found")
		return
	}
	if err != nil {
		m.internalError(w, err)
		return
	}
	if rev != cur {
		detail(w, http.StatusConflict, "Revision is not the current one")
		return
	}
	if _, err := m.deps.DB.Write.ExecContext(ctx, `INSERT INTO cc_fleet_config_acks (node_id, household_id, config_type, revision, acked_at)
		VALUES (?, ?, ?, ?, ?) ON CONFLICT (node_id, household_id, config_type) DO UPDATE SET revision = excluded.revision, acked_at = excluded.acked_at`,
		n.ID, n.HouseholdID, ctype, rev, dbTime(m.now())); err != nil {
		m.internalError(w, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"status": "acked", "revision": rev})
}
