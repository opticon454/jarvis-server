package nodesettings

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

// At-rest sealing of k2.json and node-settings.json (docs/cc/14-fleet-config-and-recovery.md
// section 8). K2 (or the household key) and every command secret live in those two files, so
// they are written as an AES-256-GCM envelope under a key that is NOT in the data directory:
//
//  1. a systemd credential, $CREDENTIALS_DIRECTORY/jarvis-at-rest-key (LoadCredentialEncrypted=
//     keeps it encrypted on disk under the TPM or the host key and hands it over per run);
//  2. the file named by $JARVIS_AT_REST_KEY_FILE;
//  3. ~/.local/state/jarvis-agent/at-rest.key (0400), made on first use.
//
// A copy of the data directory (a backup, a sync folder, a pasted listing) no longer carries
// the secrets. It does not stop someone who is root on the node, who can read the key as well.

// ErrAtRestKeyMissing: sealed files exist but no key was found. A new key is never made over
// them: it would silently turn the node's secrets into noise.
var ErrAtRestKeyMissing = errors.New("nodesettings: the node's files are encrypted but the at-rest key is missing; restore it, or reset the node")

// ErrAtRestDecrypt: the key does not open the file (a different key, another file's contents,
// or a tampered file).
var ErrAtRestDecrypt = errors.New("nodesettings: cannot decrypt the node's file with the at-rest key")

const (
	envelopeVersion = 1
	envelopeAlg     = "AES-256-GCM"
	credentialName  = "jarvis-at-rest-key"
)

type envelope struct {
	V     int    `json:"v"`
	Alg   string `json:"alg"`
	Nonce string `json:"nonce"`
	CT    string `json:"ct"` // ciphertext followed by the 16-byte tag
}

// AtRest is the key that seals the store's files and where it came from.
type AtRest struct {
	key    []byte
	Origin string
}

// NewAtRest wraps a 32-byte key. origin says where it came from, for the log and status.
func NewAtRest(key []byte, origin string) (*AtRest, error) {
	if len(key) != keyLen {
		return nil, fmt.Errorf("at-rest key must be %d bytes, got %d", keyLen, len(key))
	}
	return &AtRest{key: append([]byte(nil), key...), Origin: origin}, nil
}

// aadFor binds a ciphertext to the file it belongs to, so k2.json's contents cannot be moved
// into node-settings.json (or the reverse).
func aadFor(name string) []byte { return []byte("jarvis-at-rest:" + name) }

func (a *AtRest) seal(name string, plaintext []byte) ([]byte, error) {
	g, err := gcm(a.key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, nonceLen)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	ct := g.Seal(nil, nonce, plaintext, aadFor(name))
	return json.MarshalIndent(envelope{V: envelopeVersion, Alg: envelopeAlg, Nonce: b64(nonce), CT: b64(ct)}, "", "  ")
}

func (a *AtRest) open(name string, e envelope) ([]byte, error) {
	g, err := gcm(a.key)
	if err != nil {
		return nil, err
	}
	nonce, err1 := unb64(e.Nonce)
	ct, err2 := unb64(e.CT)
	if err := errors.Join(err1, err2); err != nil || len(nonce) != nonceLen {
		return nil, ErrAtRestDecrypt
	}
	pt, err := g.Open(nil, nonce, ct, aadFor(name))
	if err != nil {
		return nil, ErrAtRestDecrypt
	}
	return pt, nil
}

// asEnvelope reports whether b is a sealed file rather than the legacy plaintext JSON.
func asEnvelope(b []byte) (envelope, bool) {
	var e envelope
	if json.Unmarshal(b, &e) != nil || e.Alg == "" || e.CT == "" {
		return envelope{}, false
	}
	return e, true
}

func hasSealedFiles(dir string) bool {
	for _, f := range []string{K2File, SettingsFile} {
		if b, err := os.ReadFile(filepath.Join(dir, f)); err == nil {
			if _, ok := asEnvelope(b); ok {
				return true
			}
		}
	}
	return false
}

func readKeyFile(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	k, err := DecodeKey(string(trimSpace(b)))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return k, nil
}

func trimSpace(b []byte) []byte {
	for len(b) > 0 && (b[len(b)-1] == '\n' || b[len(b)-1] == '\r' || b[len(b)-1] == ' ') {
		b = b[:len(b)-1]
	}
	return b
}

// keyFile reads the key at path, making it (0400, in a 0700 directory) when it is missing and no
// sealed file depends on it.
func keyFile(path, dataDir string) ([]byte, error) {
	k, err := readKeyFile(path)
	if err == nil {
		return k, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	if hasSealedFiles(dataDir) {
		return nil, ErrAtRestKeyMissing
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return nil, err
	}
	fresh := make([]byte, keyLen)
	if _, err := rand.Read(fresh); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o400)
	if errors.Is(err, os.ErrExist) { // another process made it first
		return readKeyFile(path)
	}
	if err != nil {
		return nil, err
	}
	if _, err := f.WriteString(b64(fresh) + "\n"); err != nil {
		f.Close()
		_ = os.Remove(path)
		return nil, err
	}
	if err := f.Close(); err != nil {
		return nil, err
	}
	return fresh, nil
}

// ResolveAtRest finds the at-rest key by the order in the package comment. home is the user's
// home directory for the default key file.
func ResolveAtRest(dataDir string, env func(string) string, home string) (*AtRest, error) {
	if d := env("CREDENTIALS_DIRECTORY"); d != "" {
		if k, err := readKeyFile(filepath.Join(d, credentialName)); err == nil {
			return NewAtRest(k, "systemd credential")
		} else if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}
	path := env("JARVIS_AT_REST_KEY_FILE")
	if path == "" {
		if home == "" {
			return nil, errors.New("no home directory for the at-rest key file; set JARVIS_AT_REST_KEY_FILE")
		}
		path = filepath.Join(home, ".local", "state", "jarvis-agent", "at-rest.key")
	}
	k, err := keyFile(path, dataDir)
	if err != nil {
		return nil, err
	}
	return NewAtRest(k, "key file "+path)
}

// OpenStoreAtRest opens the store with its files sealed at rest. If no key can be had (an
// unwritable home, say) it logs a warning and opens plaintext rather than leaving the node
// without its secrets, unless JARVIS_AT_REST_REQUIRED is set. Sealed files without their key
// are always an error.
func OpenStoreAtRest(dir, nodeID string, log *slog.Logger) (*Store, error) {
	home, _ := os.UserHomeDir()
	return openStoreAtRest(dir, nodeID, log, os.Getenv, home)
}

func openStoreAtRest(dir, nodeID string, log *slog.Logger, env func(string) string, home string) (*Store, error) {
	if log == nil {
		log = slog.Default()
	}
	ar, err := ResolveAtRest(dir, env, home)
	if err != nil {
		if errors.Is(err, ErrAtRestKeyMissing) || env("JARVIS_AT_REST_REQUIRED") != "" {
			return nil, err
		}
		log.Warn("node settings: no at-rest key; K2 and secrets are stored as 0600 plaintext files", "err", err)
		st, e := OpenStoreWith(dir, nodeID, nil)
		if st != nil {
			st.atRestNote = "NOT encrypted: " + err.Error()
		}
		return st, e
	}
	st, err := OpenStoreWith(dir, nodeID, ar)
	if err == nil {
		log.Info("node settings: K2 and secrets are encrypted at rest", "key", ar.Origin)
	}
	return st, err
}
