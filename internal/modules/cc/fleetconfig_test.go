package cc

import (
	"testing"

	"github.com/alexberardi/jarvis-server/internal/platform/authn"
)

// Fleet config (docs/cc/14-fleet-config-and-recovery.md): one household-sealed blob per
// config type, pulled by every node of the household, acked per node and revision.
func TestFleetConfig(t *testing.T) {
	e := newEnv(t)
	admin := e.auth.addUser(1, "hh1", authn.RoleAdmin)
	member := e.auth.addUser(2, "hh1", authn.RoleMember)
	stranger := e.auth.addUser(3, "hh2", authn.RoleAdmin)
	n1 := e.createNode("n1", "hh1")
	n2 := e.createNode("n2", "hh1")
	foreign := e.createNode("n9", "hh2")
	c1 := e.dialNode(n1)
	c1.subscribe("jarvis/nodes/n1/#")

	body := map[string]any{"ciphertext": "c1", "nonce": "n1", "tag": "t1"}
	put := "/api/v0/households/hh1/fleet-config/secrets"

	// Writing needs household admin; another household is refused.
	e.do("PUT", put, body, bearer(member)).want(403)
	e.do("PUT", put, body, bearer(stranger)).want(403)
	e.do("PUT", put, map[string]any{"ciphertext": "c1"}, bearer(admin)).want(400)
	e.do("PUT", "/api/v0/households/hh1/fleet-config/bad%20type", body, bearer(admin)).want(400)

	r := e.do("PUT", put, body, bearer(admin)).want(200).json()
	if r["revision"].(float64) != 1 || r["nodes_notified"].(float64) != 2 {
		t.Fatalf("first put: %v", r)
	}
	// Every active household node is nudged, with no key material and no ciphertext.
	m := payload(t, c1.next())
	if m["config_type"] != "secrets" || m["revision"].(float64) != 1 || m["ciphertext"] != nil {
		t.Fatalf("nudge: %v", m)
	}

	// A member can list what the household holds (the app verifies a recovery code against it).
	l := e.do("GET", "/api/v0/households/hh1/fleet-config", nil, bearer(member)).want(200).list()
	if len(l) != 1 || l[0].(map[string]any)["ciphertext"] != "c1" {
		t.Fatalf("list: %v", l)
	}
	e.do("GET", "/api/v0/households/hh1/fleet-config", nil, bearer(stranger)).want(403)

	// Node auth is bound to the path node and its household.
	pend := "/api/v0/nodes/n1/fleet-config"
	e.do("GET", pend, nil, nil).want(400)
	e.do("GET", pend, nil, n2.h()).want(403)
	e.do("GET", pend, nil, foreign.h()).want(403)
	got := e.do("GET", pend, nil, n1.h()).want(200).list()
	if len(got) != 1 {
		t.Fatalf("pending: %v", got)
	}
	it := got[0].(map[string]any)
	if it["household_id"] != "hh1" || it["config_type"] != "secrets" || it["revision"].(float64) != 1 || it["ciphertext"] != "c1" {
		t.Fatalf("pending item: %v", it)
	}

	// Ack is per node: n2 still sees the item, n1 no longer does. A stale or future revision is refused.
	ack := pend + "/secrets/ack"
	e.do("POST", ack, map[string]any{"revision": 1}, n2.h()).want(403)
	e.do("POST", "/api/v0/nodes/n1/fleet-config/secrets/ack", map[string]any{"revision": 7}, n1.h()).want(409)
	e.do("POST", ack, map[string]any{"revision": 1}, n1.h()).want(200)
	if got := e.do("GET", pend, nil, n1.h()).want(200).list(); len(got) != 0 {
		t.Fatalf("acked item still pending: %v", got)
	}
	if got := e.do("GET", "/api/v0/nodes/n2/fleet-config", nil, n2.h()).want(200).list(); len(got) != 1 {
		t.Fatalf("n2 must still see it: %v", got)
	}

	// A new revision supersedes the old one and is pending again for n1.
	body2 := map[string]any{"ciphertext": "c2", "nonce": "n2", "tag": "t2"}
	if r := e.do("PUT", put, body2, bearer(admin)).want(200).json(); r["revision"].(float64) != 2 {
		t.Fatalf("second put: %v", r)
	}
	got = e.do("GET", pend, nil, n1.h()).want(200).list()
	if len(got) != 1 || got[0].(map[string]any)["ciphertext"] != "c2" {
		t.Fatalf("revision 2 not pending: %v", got)
	}
	// The ack of a superseded revision does not hide the new one.
	e.do("POST", ack, map[string]any{"revision": 1}, n1.h()).want(409)

	// A node added after the first push gets it on its first pull.
	late := e.createNode("n3", "hh1")
	if got := e.do("GET", "/api/v0/nodes/n3/fleet-config", nil, late.h()).want(200).list(); len(got) != 1 {
		t.Fatalf("late joiner: %v", got)
	}

	// Removing a config type deletes the ciphertext (rotate or clear).
	e.do("DELETE", put, nil, bearer(member)).want(403)
	e.do("DELETE", put, nil, bearer(admin)).want(200)
	if l := e.do("GET", "/api/v0/households/hh1/fleet-config", nil, bearer(admin)).want(200).list(); len(l) != 0 {
		t.Fatalf("after delete: %v", l)
	}
}
