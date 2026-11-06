//go:build contract

package contract

import (
	"encoding/json"
	"net/http"
	"testing"
)

// Fleet config is jarvisd-only (docs/cc/14-fleet-config-and-recovery.md): a household-sealed
// blob per config type that every node of the household pulls and acks. The shapes below are
// what the mobile app and the Go node depend on.
func TestCCFleetConfig(t *testing.T) {
	tg := T(t)
	if !Jarvisd() {
		t.Skip("contract: fleet config is jarvisd-only (no legacy equivalent); set " + EnvImpl + "=jarvisd")
	}
	tg.NeedCCAdmin(t)
	tg.NeedMQTT(t)
	u := NewUser(t)
	n := NewCCNode(t, u)
	stranger := NewUser(t)
	hh := u.HouseholdID
	put := "/api/v0/households/" + hh + "/fleet-config/secrets"
	pend := "/api/v0/nodes/" + n.ID + "/fleet-config"

	c := DialMQTTNode(t, n)
	c.Subscribe(t, "jarvis/nodes/"+n.ID+"/#")

	body := map[string]any{"ciphertext": "Y29udHJhY3Q", "nonce": "bm9uY2U", "tag": "dGFn"}
	tg.Do(t, CommandCenter, http.MethodPut, put, body, stranger.H()).ExpectStatus(http.StatusForbidden)
	tg.Do(t, CommandCenter, http.MethodPut, put, map[string]any{"ciphertext": "x"}, u.H()).ExpectStatus(http.StatusBadRequest)
	tg.Do(t, CommandCenter, http.MethodPut, put, body, u.H()).Expect(http.StatusOK, Obj{
		"config_type": Eq("secrets"), "revision": Int, "nodes_notified": Int,
	})
	// The nudge carries ids only; the node pulls the ciphertext over its authenticated channel.
	ExpectMQTT(t, c.NextTopic(t, mqttWait, "jarvis/nodes/"+n.ID+"/fleet/config"),
		Obj{"household_id": Eq(hh), "config_type": Eq("secrets"), "revision": Int})

	item := Obj{"household_id": Eq(hh), "config_type": Eq("secrets"), "ciphertext": String, "nonce": String, "tag": String, "revision": Int}
	tg.Get(t, CommandCenter, "/api/v0/households/"+hh+"/fleet-config", u.H()).Expect(http.StatusOK, ArrayOf(item))
	tg.Get(t, CommandCenter, "/api/v0/households/"+hh+"/fleet-config", stranger.H()).ExpectStatus(http.StatusForbidden)

	tg.Get(t, CommandCenter, pend, n.APIKeyH()).Expect(http.StatusOK, NonEmptyArrayOf(item))
	rev, _ := tg.Get(t, CommandCenter, pend, n.APIKeyH()).JSON().([]any)[0].(map[string]any)["revision"].(json.Number).Int64()
	ack := pend + "/secrets/ack"
	tg.Post(t, CommandCenter, ack, map[string]any{"revision": rev + 5}, n.APIKeyH()).ExpectStatus(http.StatusConflict)
	tg.Post(t, CommandCenter, ack, map[string]any{"revision": rev}, n.APIKeyH()).Expect(http.StatusOK, Obj{"status": Eq("acked"), "revision": Int})
	tg.Get(t, CommandCenter, pend, n.APIKeyH()).Expect(http.StatusOK, EqDeep([]any{}))

	tg.Do(t, CommandCenter, http.MethodDelete, put, nil, stranger.H()).ExpectStatus(http.StatusForbidden)
	tg.Do(t, CommandCenter, http.MethodDelete, put, nil, u.H()).Expect(http.StatusOK, Obj{"status": Eq("deleted")})
}
