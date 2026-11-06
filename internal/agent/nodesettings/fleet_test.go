package nodesettings

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
)

func sealedFleet(t *testing.T, key []byte, hh, ctype string, rev int, v any) map[string]any {
	t.Helper()
	pt, _ := json.Marshal(v)
	s, err := SealFleet(key, hh, ctype, pt, nil)
	if err != nil {
		t.Fatal(err)
	}
	return map[string]any{"household_id": hh, "config_type": ctype, "revision": rev,
		"ciphertext": s.Ciphertext, "nonce": s.Nonce, "tag": s.Tag}
}

func TestFleetAADBindsHouseholdTypeAndKind(t *testing.T) {
	key, _ := DecodeKey(testKey)
	s, err := SealFleet(key, "hh1", "secrets", []byte(`{"A":"1"}`), nil)
	if err != nil {
		t.Fatal(err)
	}
	if pt, err := OpenFleet(key, "hh1", "secrets", s); err != nil || string(pt) != `{"A":"1"}` {
		t.Fatalf("round trip: %q %v", pt, err)
	}
	if _, err := OpenFleet(key, "hh2", "secrets", s); !errors.Is(err, ErrDecrypt) {
		t.Errorf("another household opened it: %v", err)
	}
	if _, err := OpenFleet(key, "hh1", "node_settings", s); !errors.Is(err, ErrDecrypt) {
		t.Errorf("another type opened it: %v", err)
	}
	// A per-node message is not fleet config and the reverse: the two AADs are different strings.
	if _, err := Open(key, "hh1", "secrets", s); !errors.Is(err, ErrDecrypt) {
		t.Errorf("fleet message opened as a node message: %v", err)
	}
	n, _ := Seal(key, "n1", "secrets", []byte(`{}`), nil)
	if _, err := OpenFleet(key, "n1", "secrets", n); !errors.Is(err, ErrDecrypt) {
		t.Errorf("node message opened as fleet config: %v", err)
	}
	other, _ := DecodeKey(b64([]byte("abcdef0123456789abcdef0123456789")))
	if _, err := OpenFleet(other, "hh1", "secrets", s); !errors.Is(err, ErrDecrypt) {
		t.Errorf("a different household key opened it: %v", err)
	}
}

func TestFleetConfigAppliesAndAcks(t *testing.T) {
	h, f, _ := newHandler(t, "n1")
	if err := h.Store.SaveK2(testKey, "kid", ""); err != nil {
		t.Fatal(err)
	}
	key := h.Store.Key()
	var gotSettings map[string]any
	h.NodeSettings = func(m map[string]any) error { gotSettings = m; return nil }
	changed := 0
	h.SecretsChanged = func() { changed++ }

	other, _ := DecodeKey(b64([]byte("abcdef0123456789abcdef0123456789")))
	f.fleet = []map[string]any{
		sealedFleet(t, key, "hh1", "secrets", 3, map[string]any{"WEATHER_API_KEY": "abc", "UNITS": "metric"}),
		sealedFleet(t, key, "hh1", "command_registry", 1, map[string]any{"command_name": "get_weather", "enabled": false}),
		sealedFleet(t, key, "hh1", "node_settings", 2, map[string]any{"volume_percent": 40}),
		sealedFleet(t, other, "hh1", "secrets2", 1, map[string]any{"X": "y"}), // someone else's key
	}
	h.Handle(context.Background(), TopicFleetConfig, []byte(`{"household_id":"hh1","config_type":"secrets","revision":3}`))

	if h.Store.Secret("WEATHER_API_KEY", "integration", nil) != "abc" || h.Store.Secret("UNITS", "integration", nil) != "metric" {
		t.Error("fleet secrets not stored")
	}
	if h.Store.Enabled("get_weather") {
		t.Error("fleet command_registry did not disable get_weather")
	}
	if gotSettings["volume_percent"] != json.Number("40") {
		t.Errorf("node settings hook got %v", gotSettings)
	}
	if changed != 1 {
		t.Errorf("SecretsChanged ran %d times, want 1", changed)
	}
	base := "POST /api/v0/nodes/n1/fleet-config/"
	for ct, rev := range map[string]float64{"secrets": 3, "command_registry": 1, "node_settings": 2} {
		if !f.called(base + ct + "/ack") {
			t.Errorf("%s not acked", ct)
		} else if got := f.body(base + ct + "/ack")["revision"]; got != rev {
			t.Errorf("%s acked revision %v, want %v", ct, got, rev)
		}
	}
	if f.called(base + "secrets2/ack") {
		t.Error("a message sealed with another key was acked")
	}
	// Nothing was written for the failed item.
	if h.Store.Secret("X", "integration", nil) != "" {
		t.Error("an undecryptable item was applied")
	}
}

func TestFleetConfigWaitsForTheKey(t *testing.T) {
	h, f, _ := newHandler(t, "n1")
	key, _ := DecodeKey(testKey)
	f.fleet = []map[string]any{sealedFleet(t, key, "hh1", "secrets", 1, map[string]any{"A": "1"})}
	if n := h.FleetConfig(context.Background()); n != 0 {
		t.Fatalf("applied %d with no K2", n)
	}
	if f.called("POST /api/v0/nodes/n1/fleet-config/secrets/ack") {
		t.Fatal("acked without a key")
	}
	// Once the key arrives (a node added after fleet config exists) the next pull applies it.
	if err := h.Store.SaveK2(testKey, "kid", ""); err != nil {
		t.Fatal(err)
	}
	if n := h.FleetConfig(context.Background()); n != 1 || h.Store.Secret("A", "integration", nil) != "1" {
		t.Fatalf("late key: applied %d", n)
	}
}

func TestFleetNodeSettingsWithoutAHookIsAckedNotApplied(t *testing.T) {
	h, f, _ := newHandler(t, "n1")
	_ = h.Store.SaveK2(testKey, "kid", "")
	f.fleet = []map[string]any{sealedFleet(t, h.Store.Key(), "hh1", "node_settings", 1, map[string]any{"volume_percent": 40})}
	h.FleetConfig(context.Background())
	if !f.called("POST /api/v0/nodes/n1/fleet-config/node_settings/ack") {
		t.Error("a computer must ack node settings it has nothing to apply, not leave them pending forever")
	}
}

func TestBackstopPullsFleetConfig(t *testing.T) {
	h, f, _ := newHandler(t, "n1")
	_ = h.Store.SaveK2(testKey, "kid", "")
	f.fleet = []map[string]any{sealedFleet(t, h.Store.Key(), "hh1", "secrets", 1, map[string]any{"B": "2"})}
	h.Backstop(context.Background())
	if h.Store.Secret("B", "integration", nil) != "2" {
		t.Error("backstop did not pull fleet config")
	}
}

// A node added after fleet config exists: the household key arrives through the K2 relay and
// the node pulls and applies everything the household already holds, without a reconnect.
func TestK2ProvisionPullsExistingFleetConfig(t *testing.T) {
	h, f, _ := newHandler(t, "n1")
	key, _ := DecodeKey(testKey)
	f.k2 = map[string]any{"k2": testKey, "kid": "k2-2026-11", "created_at": "2026-11-06T00:00:00Z"}
	f.fleet = []map[string]any{sealedFleet(t, key, "hh1", "secrets", 4, map[string]any{"LATE": "joiner"})}
	h.Handle(context.Background(), TopicK2Provision, []byte(`{"request_id":"r1"}`))
	if h.Store.Secret("LATE", "integration", nil) != "joiner" {
		t.Fatal("a new node did not pick up the household's fleet config after receiving the key")
	}
	if got := f.body("POST /api/v0/nodes/n1/fleet-config/secrets/ack")["revision"]; got != float64(4) {
		t.Errorf("acked revision %v, want 4", got)
	}
}

func TestFleetCommandStatesSwitchSeveralCommands(t *testing.T) {
	h, f, _ := newHandler(t, "n1")
	_ = h.Store.SaveK2(testKey, "kid", "")
	f.fleet = []map[string]any{sealedFleet(t, h.Store.Key(), "hh1", "command_states", 1,
		map[string]any{"get_weather": false, "tell_joke": true})}
	h.FleetConfig(context.Background())
	if h.Store.Enabled("get_weather") || !h.Store.Enabled("tell_joke") {
		t.Errorf("command states: weather %v, joke %v", h.Store.Enabled("get_weather"), h.Store.Enabled("tell_joke"))
	}
}

// A vector made by Node's crypto (AES-256-GCM, the mobile app's side of the wire) pins the fleet
// AAD string "{household}:fleet:{type}" across the two implementations. The same vector is in
// the app's tests (__tests__/services/fleetConfigService.test.ts).
func TestFleetSealMatchesTheApp(t *testing.T) {
	key, err := DecodeKey("MDEyMzQ1Njc4OWFiY2RlZjAxMjM0NTY3ODlhYmNkZWY")
	if err != nil {
		t.Fatal(err)
	}
	want := Sealed{Ciphertext: "Eu6e-oqxJZYxQ7NBQ8YTegUTgT3Y4gx5Yw", Nonce: "YWJjZGVmZ2hpamts", Tag: "GA7I3CryDzXorVsoTisqEg"}
	got, err := SealFleet(key, "hh1", "secrets", []byte(`{"WEATHER_API_KEY":"abc"}`), []byte("abcdefghijkl"))
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("SealFleet = %+v, want %+v", got, want)
	}
	pt, err := OpenFleet(key, "hh1", "secrets", want)
	if err != nil || string(pt) != `{"WEATHER_API_KEY":"abc"}` {
		t.Fatalf("OpenFleet = %q, %v", pt, err)
	}
}

// Removing a value from the fleet removes it from every node, and only what the fleet put there.
func TestFleetSecretsRemovalReachesTheNode(t *testing.T) {
	h, f, _ := newHandler(t, "n1")
	_ = h.Store.SaveK2(testKey, "kid", "")
	key := h.Store.Key()
	_ = h.Store.SetSecret("PER_NODE", "mine", "integration", "", nil) // set on this node alone
	ctx := context.Background()

	f.fleet = []map[string]any{sealedFleet(t, key, "hh1", "secrets", 1, map[string]any{"A": "1", "B": "2"})}
	h.FleetConfig(ctx)
	if h.Store.Secret("A", "integration", nil) != "1" || h.Store.Secret("B", "integration", nil) != "2" {
		t.Fatal("setup: fleet secrets not applied")
	}

	f.fleet = []map[string]any{sealedFleet(t, key, "hh1", "secrets", 2, map[string]any{"A": "1"})}
	h.FleetConfig(ctx)
	if h.Store.Secret("B", "integration", nil) != "" {
		t.Error("B was removed from the fleet but is still on the node")
	}
	if h.Store.Secret("A", "integration", nil) != "1" {
		t.Error("A was removed although the fleet still has it")
	}
	if h.Store.Secret("PER_NODE", "integration", nil) != "mine" {
		t.Error("a secret the fleet never set was deleted")
	}

	// The record of what the fleet set survives a restart, so a removal after a reboot still works.
	again, err := OpenStore(h.Store.dir, "n1")
	if err != nil {
		t.Fatal(err)
	}
	h2 := &Handler{NodeID: "n1", API: f, Store: again, Log: quiet()}
	f.fleet = []map[string]any{sealedFleet(t, again.Key(), "hh1", "secrets", 3, map[string]any{})}
	h2.FleetConfig(ctx)
	if again.Secret("A", "integration", nil) != "" {
		t.Error("A survived the fleet emptying the type after a restart")
	}
	if again.Secret("PER_NODE", "integration", nil) != "mine" {
		t.Error("a per-node secret was deleted by the fleet")
	}
}
