package nodesettings

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func testAtRest(t *testing.T, fill byte) *AtRest {
	t.Helper()
	ar, err := NewAtRest(bytes.Repeat([]byte{fill}, 32), "test")
	if err != nil {
		t.Fatal(err)
	}
	return ar
}

func readAll(t *testing.T, dir string) string {
	t.Helper()
	var sb strings.Builder
	for _, f := range []string{K2File, SettingsFile} {
		if b, err := os.ReadFile(filepath.Join(dir, f)); err == nil {
			sb.Write(b)
		}
	}
	return sb.String()
}

func TestAtRestFilesHoldNoPlaintextAndReload(t *testing.T) {
	dir := t.TempDir()
	st, err := OpenStoreWith(dir, "n1", testAtRest(t, 1))
	if err != nil {
		t.Fatal(err)
	}
	if err := st.SaveK2(testKey, "kid-1", "2026-11-06T00:00:00Z"); err != nil {
		t.Fatal(err)
	}
	if err := st.SetSecret("WEATHER_API_KEY", "s3cr3t-value", "integration", "", nil); err != nil {
		t.Fatal(err)
	}
	raw := readAll(t, dir)
	for _, leak := range []string{testKey, "s3cr3t-value", "WEATHER_API_KEY", "kid-1"} {
		if strings.Contains(raw, leak) {
			t.Errorf("on-disk files contain %q", leak)
		}
	}
	if st.AtRest() != "encrypted (test)" {
		t.Errorf("AtRest() = %q", st.AtRest())
	}
	again, err := OpenStoreWith(dir, "n1", testAtRest(t, 1))
	if err != nil {
		t.Fatal(err)
	}
	if !again.HasK2() || again.Secret("WEATHER_API_KEY", "integration", nil) != "s3cr3t-value" {
		t.Fatal("a reopened store lost its K2 or secret")
	}
	if runtime.GOOS != "windows" {
		for _, f := range []string{K2File, SettingsFile} {
			if fi, _ := os.Stat(filepath.Join(dir, f)); fi.Mode().Perm() != 0o600 {
				t.Errorf("%s mode %o", f, fi.Mode().Perm())
			}
		}
	}
}

func TestAtRestWrongKeyAndMissingKeyAreErrorsNotDataLoss(t *testing.T) {
	dir := t.TempDir()
	st, _ := OpenStoreWith(dir, "n1", testAtRest(t, 1))
	_ = st.SaveK2(testKey, "kid", "")
	before := readAll(t, dir)

	if _, err := OpenStoreWith(dir, "n1", testAtRest(t, 2)); !errors.Is(err, ErrAtRestDecrypt) {
		t.Fatalf("wrong key: %v", err)
	}
	if _, err := OpenStoreWith(dir, "n1", nil); !errors.Is(err, ErrAtRestKeyMissing) {
		t.Fatalf("no key: %v", err)
	}
	if readAll(t, dir) != before {
		t.Fatal("a failed open changed the files")
	}
}

func TestAtRestMigratesPlaintextOnFirstOpen(t *testing.T) {
	dir := t.TempDir()
	plain, _ := OpenStore(dir, "n1")
	_ = plain.SaveK2(testKey, "kid", "")
	_ = plain.SetSecret("A", "legacy-secret", "integration", "", nil)
	if !strings.Contains(readAll(t, dir), "legacy-secret") {
		t.Fatal("setup: expected plaintext")
	}
	st, err := OpenStoreWith(dir, "n1", testAtRest(t, 1))
	if err != nil {
		t.Fatal(err)
	}
	if !st.HasK2() || st.Secret("A", "integration", nil) != "legacy-secret" {
		t.Fatal("migration lost data")
	}
	if raw := readAll(t, dir); strings.Contains(raw, "legacy-secret") || strings.Contains(raw, testKey) {
		t.Error("files are still plaintext after the first sealed open")
	}
}

func TestAtRestBindsCiphertextToItsFile(t *testing.T) {
	dir := t.TempDir()
	st, _ := OpenStoreWith(dir, "n1", testAtRest(t, 1))
	_ = st.SaveK2(testKey, "kid", "")
	_ = st.SetSecret("A", "v", "integration", "", nil)
	k2, _ := os.ReadFile(filepath.Join(dir, K2File))
	if err := os.WriteFile(filepath.Join(dir, SettingsFile), k2, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenStoreWith(dir, "n1", testAtRest(t, 1)); !errors.Is(err, ErrAtRestDecrypt) {
		t.Fatalf("a file swapped for another opened: %v", err)
	}
}

func envOf(m map[string]string) func(string) string { return func(k string) string { return m[k] } }

func TestResolveAtRestPrefersASystemdCredentialThenAnExplicitFile(t *testing.T) {
	dir, home := t.TempDir(), t.TempDir()
	cred := t.TempDir()
	key1 := b64(bytes.Repeat([]byte{7}, 32))
	if err := os.WriteFile(filepath.Join(cred, "jarvis-at-rest-key"), []byte(key1+"\n"), 0o400); err != nil {
		t.Fatal(err)
	}
	explicit := filepath.Join(t.TempDir(), "k")
	_ = os.WriteFile(explicit, []byte(b64(bytes.Repeat([]byte{8}, 32))), 0o400)

	ar, err := ResolveAtRest(dir, envOf(map[string]string{"CREDENTIALS_DIRECTORY": cred, "JARVIS_AT_REST_KEY_FILE": explicit}), home)
	if err != nil || ar.Origin != "systemd credential" || !bytes.Equal(ar.key, bytes.Repeat([]byte{7}, 32)) {
		t.Fatalf("credential first: %v %+v", err, ar)
	}
	ar, err = ResolveAtRest(dir, envOf(map[string]string{"JARVIS_AT_REST_KEY_FILE": explicit}), home)
	if err != nil || !strings.HasPrefix(ar.Origin, "key file ") || !bytes.Equal(ar.key, bytes.Repeat([]byte{8}, 32)) {
		t.Fatalf("explicit file: %v %+v", err, ar)
	}
	if _, err := os.Stat(filepath.Join(home, ".local", "state", "jarvis-agent", "at-rest.key")); err == nil {
		t.Error("a default key was made although another source was configured")
	}
}

func TestResolveAtRestCreatesOneKeyOutsideTheDataDirAndKeepsIt(t *testing.T) {
	dir, home := t.TempDir(), t.TempDir()
	ar1, err := ResolveAtRest(dir, envOf(nil), home)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(home, ".local", "state", "jarvis-agent", "at-rest.key")
	fi, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if runtime.GOOS != "windows" && fi.Mode().Perm() != 0o400 {
		t.Errorf("key file mode %o, want 0400", fi.Mode().Perm())
	}
	if strings.HasPrefix(path, dir) {
		t.Error("the key file is inside the data dir")
	}
	ar2, err := ResolveAtRest(dir, envOf(nil), home)
	if err != nil || !bytes.Equal(ar1.key, ar2.key) {
		t.Fatalf("second resolve changed the key: %v", err)
	}
}

func TestResolveAtRestNeverMakesANewKeyOverExistingCiphertext(t *testing.T) {
	dir, home := t.TempDir(), t.TempDir()
	st, _ := OpenStoreWith(dir, "n1", testAtRest(t, 1))
	_ = st.SaveK2(testKey, "kid", "")
	// The key is gone (a restored backup without it): a silent new key would make the data unreadable.
	if _, err := ResolveAtRest(dir, envOf(nil), home); !errors.Is(err, ErrAtRestKeyMissing) {
		t.Fatalf("got %v, want ErrAtRestKeyMissing", err)
	}
	if _, err := os.Stat(filepath.Join(home, ".local", "state", "jarvis-agent", "at-rest.key")); err == nil {
		t.Error("a replacement key was created")
	}
}

func TestOpenStoreAtRestFallsBackToPlaintextWithAWarningUnlessRequired(t *testing.T) {
	dir := t.TempDir()
	noHome := t.TempDir()
	_ = os.Chmod(noHome, 0o500) // the key cannot be created
	if runtime.GOOS == "windows" || os.Geteuid() == 0 {
		t.Skip("directory permissions are not enforced here")
	}
	st, err := openStoreAtRest(dir, "n1", quiet(), envOf(nil), noHome)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(st.AtRest(), "NOT encrypted") {
		t.Errorf("AtRest() = %q", st.AtRest())
	}
	if _, err := openStoreAtRest(dir, "n1", quiet(), envOf(map[string]string{"JARVIS_AT_REST_REQUIRED": "1"}), noHome); err == nil {
		t.Error("JARVIS_AT_REST_REQUIRED=1 must refuse to run unencrypted")
	}
}
