# 14. Fleet config and recovery

Set a secret or a node setting once, and every node in the household (including nodes added later) receives it. The key that protects it stays on your phone and your nodes, never on the server, and a recovery code lets you get back in if the phone is lost.

**Status:** built on branch `feature/fleet-node-config`. Server and Go node: this repo. App: `jarvis-node-mobile` (same branch name). Python nodes are out of scope.

## 1. The problem this solves

Before this, each node had its own key (K2). Every secret or setting was sealed for one node and pushed to that node from the app. That meant:

- Entering the same API key once per node.
- A node added later missed everything pushed before it existed.
- Losing the phone lost every K2, so every node had to be re-keyed by hand.

## 2. How it works

```text
phone                         jarvisd (ciphertext only)                node
-----                         -------------------------                ----
seal config with KH  ──PUT──▶  cc_fleet_config (household, type)
                               revision + 1
                               nudge ──── MQTT fleet/config ─────────▶ GET /nodes/{id}/fleet-config
                                                                       open with KH, apply
                               cc_fleet_config_acks  ◀── POST ack ──── ack {revision}
```

- **KH (household key).** One random 256-bit key per household. The phone creates it. Each node stores it in the slot a per-node K2 used, so the existing K2 relay delivers it (`POST /nodes/{id}/k2`).
- **Sealing.** AES-256-GCM, a fresh 12-byte nonce, and the additional data `{household_id}:fleet:{config_type}`. The data is bound to a household and a type, not a node, so one stored ciphertext serves every node. A per-node message (`{node_id}:{config_type}`) never opens as fleet config, and the reverse.
- **The server.** jarvisd stores `ciphertext`, `nonce`, `tag` and a `revision`. It cannot read them. The MQTT nudge carries ids only.
- **Applying once.** The server offers each revision to each node until that node acks it. A node applies a revision once, so a per-node override pushed afterwards stays in force until the next fleet change to the same type.

### Config types

| Type             | Plaintext (JSON object)                       | Effect on a node                                                             |
| ---------------- | --------------------------------------------- | ---------------------------------------------------------------------------- |
| `secrets`        | `{ "NAME": "value", ... }`                    | Stored as integration secrets, as a per-node secrets push does.              |
| `node_settings`  | `{ "volume_percent": 40, "is_muted": false }` | Applied like `update_node_config` (live, then saved). Trust anchors dropped. |
| `command_states` | `{ "get_weather": false, "tell_joke": true }` | Turns each command on or off.                                                |

Any other type name is stored and offered, and a node treats it as `secrets`. Use the three above.

## 3. Routes

All under `/api/v0` on the command-centre listener. Errors follow the rest of cc: a missing field is a `400` validation error.

| Method | Path                                   | Caller           | Notes                                                                               |
| ------ | -------------------------------------- | ---------------- | ----------------------------------------------------------------------------------- |
| PUT    | `/households/{hh}/fleet-config/{type}` | household admin  | Body `{ciphertext, nonce, tag}`. Returns `{config_type, revision, nodes_notified}`. |
| DELETE | `/households/{hh}/fleet-config/{type}` | household admin  | Removes the blob and its acks.                                                      |
| GET    | `/households/{hh}/fleet-config`        | household member | Lists the sealed items. The app uses it to check a recovery code.                   |
| GET    | `/nodes/{id}/fleet-config`             | the node         | Items this node has not applied. Bound to the path node (D6).                       |
| POST   | `/nodes/{id}/fleet-config/{type}/ack`  | the node         | Body `{revision}`. `409` unless it is the current revision.                         |

Type names match `^[a-z0-9][a-z0-9_.:-]{0,63}$`. The tables are `cc_fleet_config` and `cc_fleet_config_acks` (cc migration 00350).

## 4. The recovery code

The code is the household key written for a person: 60 characters of Crockford base32 in twelve groups of five.

```text
04002-0G30G-2GC1R-81450-P30D1-R7H04-8J2CA-1A5GQ-30CHM-6RW3M-F1ZH5-2G5CG
```

That example is the key made of bytes 0 to 31, and the app's tests pin it.

- **Layout.** One version byte (`0x01`), the 32 key bytes, then a CRC-32 of those 33 bytes (big endian). 37 bytes become 296 bits, padded with four zero bits to 300 bits, which is 60 characters.
- **Typing it back.** Lower case, spaces, missing dashes, `O` for `0`, and `I` or `L` for `1` are all accepted. `U` is not used.
- **What the CRC does.** It catches typos. It does not prove the code is the right key. Only opening something sealed with the key does that (section 6).
- **Treat it like a BitLocker recovery key.** Anyone holding the code can read the fleet secrets. There is no escrow and no server copy. Lose the phone and the code, and the fleet secrets are unrecoverable: you create a new key and enter them again.

## 5. Day to day

### First time

1. In the app, open **Settings, All nodes, Recovery code** and choose **Create household key**.
2. The code appears once. Copy it into your password manager (1Password and similar) before you leave the screen. Choose **Hide**.
3. Choose **Give key to all nodes**. Each node receives the key through the K2 relay and then pulls the fleet settings. A node that is offline fails with a clear message; run it again when it is back.
4. Open **Settings, All nodes, Fleet settings**. Pick **Secrets**, **Node** or **Commands**, enter a name and value, and choose **Save for all nodes**. The screen reports how many nodes were told. Offline nodes catch up when they reconnect.

### Adding a node later

Set the node up as usual, then on the Recovery code screen choose **Give key to all nodes** (it is safe to repeat). The node pulls everything the household already holds the moment it has the key, with no further steps. You do not re-enter any secret.

### Per-node overrides

Pushing a value to one node (the node's own settings screen) still works. A node applies each fleet revision once, so the override stays until you change that fleet type again. To make a node follow the fleet again, set the value in Fleet settings.

## 6. If the phone is lost, wiped or replaced

You need the recovery code and your jarvisd login.

1. Install the app on the new phone and sign in.
2. Open **Settings, All nodes, Recovery code**. Under **Restore on this phone**, paste or type the code and choose **Restore**.
3. The app lists what the server holds and tries to open it with the code.
   - **"Restored. The code opened this household's stored settings"**: the code is proven right.
   - **"That code is well formed but it is not this household's key"**: wrong code (maybe another household's). Nothing was changed on the phone.
   - **"The code has a typo"**: re-check the characters.
   - **"Restored, but nothing is stored ... could not be checked"**: the household has no fleet settings yet, so there was nothing to test the code against. Set one value in Fleet settings; if it will not open later, the code was wrong.
4. If the old phone was stolen rather than lost, rotate the key (section 7) as soon as you have restored.

With no recovery code the old key is gone. Create a new household key, choose **Give key to all nodes**, and enter the fleet secrets again. The old ciphertext on the server is useless and is replaced as you save new values.

## 7. Rotating the key

Use **Rotate key** when the code leaked, a phone was stolen, or someone left the household.

1. The app opens every stored item with the current key (if any cannot be opened, nothing changes).
2. It makes a new key, keeping the old one on the phone as `previous` until the rotation finishes.
3. It seals every item with the new key and writes it back.
4. It gives the new key to every node through the K2 relay. Each node then pulls and applies the resealed items.
5. It deletes `previous` and shows the new recovery code. Save it in your vault and delete the old entry.

If a node is offline the rotation stops with an error naming it. Run **Rotate key** again once it is online: the app resumes with the same new key. Until a node has the new key it cannot open new items and keeps its last applied settings.

## 8. What each party can see

| Party                      | Can read fleet secrets? | Notes                                                                 |
| -------------------------- | ----------------------- | --------------------------------------------------------------------- |
| jarvisd and its database   | No                      | Ciphertext, nonce, tag, revision and the household id only.           |
| A backup of jarvisd        | No                      | Same.                                                                 |
| MQTT broker and observers  | No                      | The nudge has ids only.                                               |
| A household member's phone | Only with the key       | The key is in the phone's secure store, scoped to the signed-in user. |
| A node                     | Yes                     | It must, to apply them. See the limits below.                         |
| Anyone with the code       | Yes                     | Rotate if it may have leaked.                                         |

### Known limits

- **Root on a node can still read them.** The node seals `k2.json` and `node-settings.json` (section 9), but the node must be able to use the key, so someone who is root on the node can read the at-rest key too. The seal protects copies of the data directory, not a compromised node.
- **Rotation needs the nodes online.** A node that was offline during a rotation cannot open the new items until it receives the new key.
- **Command states are replaced as a whole.** `command_states` is one blob, so Fleet settings reads the current values and merges your change. Two phones editing at once can overwrite each other.
- **Python nodes** do not understand fleet config. They ignore the nudge.

## 9. Troubleshooting

| Symptom                                                | Likely cause and fix                                                                                                     |
| ------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| "Give key to all nodes" says it could not reach a node | The node is offline or not on MQTT. Check it, then run again.                                                            |
| A node did not apply a value                           | Check its log for `fleet:`. "no household key is provisioned" means it never received KH: run **Give key to all nodes**. |
| `fleet: config not applied ... decryption failed`      | The node holds a different key (an old per-node K2, or a half-finished rotation). Give it the current key.               |
| Value applied, then a per-node value took over         | Expected: per-node pushes win until the next fleet change to that type.                                                  |
| Fleet settings says "No household key on this phone"   | Create the key, or restore it from the recovery code.                                                                    |
| Restore says "not this household's key"                | The code is from another household or an older key. Find the newest entry in your vault, or rotate after restoring.      |

## 9. Encryption at rest on a node

`k2.json` (the household key) and `node-settings.json` (secrets, command switches and command data) are written as an AES-256-GCM envelope, not as readable JSON. The envelope is bound to its file name, so one file's contents cannot be moved into the other.

The key that seals them is not in the data directory. The node looks for it in this order:

1. **A systemd credential**, `$CREDENTIALS_DIRECTORY/jarvis-at-rest-key`. Prefer this on a Pi or server: `LoadCredentialEncrypted=` keeps the key encrypted on disk under the TPM or the host key and hands it to the service for each run.
2. **A key file** named by `JARVIS_AT_REST_KEY_FILE`.
3. **The default key file**, `~/.local/state/jarvis-agent/at-rest.key` (mode 0400, directory 0700), made on first start. This is a different directory from the node's config, so copying `~/.config/jarvis-agent` does not copy the key.

What happens in the awkward cases:

| Case                                               | Result                                                                                                         |
| -------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| Existing plaintext files from before this change   | Read once, then rewritten sealed on that first start.                                                          |
| Sealed files but the key is missing                | The node refuses to open the store (`at-rest key is missing`). It never makes a new key over them.             |
| The key is wrong, or a file was altered or swapped | The store refuses to open (`cannot decrypt`). Nothing is overwritten.                                          |
| No key can be made (read-only home, no home)       | The node logs a warning and keeps plaintext 0600 files so it still works. `JARVIS_AT_REST_REQUIRED=1` refuses. |

The start-up log says which applies (`K2 and secrets are encrypted at rest key=...`, or the warning).

To use a systemd credential on a node:

```bash
sudo mkdir -p /etc/jarvis-node
head -c 32 /dev/urandom | basenc --base64url -w0 | sudo systemd-creds encrypt --name=jarvis-at-rest-key - /etc/jarvis-node/at-rest.cred
# in the unit: LoadCredentialEncrypted=jarvis-at-rest-key:/etc/jarvis-node/at-rest.cred
```

Back up the key (or the credential and the host key) with the node's backups, or a restored backup will not open. A node you cannot recover is reset and re-paired: the secrets then come back from the fleet settings.

Not covered: the node's own login (`node.json`, which holds the node key) is still a 0600 plaintext file.

## 10. Where the code is

| Area                    | Files                                                                                                                                  |
| ----------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| Server routes and table | `internal/modules/cc/fleetconfig.go`, `internal/modules/cc/migrations/00350_fleet_config.sql`                                          |
| Node pull and apply     | `internal/agent/nodesettings/handler.go` (`FleetConfig`), `crypto.go` (`SealFleet`, `OpenFleet`)                                       |
| Node settings hook      | `internal/agent/pinode/pinode.go` (`ApplySettings`), wired in `cmd/jarvis-agent/main.go`                                               |
| Contract test           | `contract/cc_fleet_config_test.go` (jarvisd-only)                                                                                      |
| App                     | `jarvis-node-mobile`: `recoveryCodeService`, `householdKeyService`, `fleetConfigService`, the Recovery code and Fleet settings screens |

## 11. How it is tested

- **Unit and contract tests** cover the sealing, the routes, the node's pull and apply, at-rest sealing and the app's services and screens.
- **`scripts/e2e-fleet-config`** runs the app's real services (under Jest in `jarvis-node-mobile`) against a real jarvisd and real `jarvis-agent` nodes in a sandbox. It creates the key, pushes secrets, node settings and command states, and gives node A the key. It checks the node's log and opens the node's sealed `node-settings.json` with its at-rest key. It then wipes the app's storage, restores from the typed recovery code (lower case, spaces for dashes), confirms a wrong code is refused, and pushes a change that node A applies. Finally a new phone holding only the code adds node B after the key exists, and B receives everything with no secret entered again. Run it from a disk-backed copy of the repo, because SQLite does not work on an SMB share.
- **Not covered:** taps through the app screens on a device (no Maestro flows yet), and real Pi or Surface hardware.
