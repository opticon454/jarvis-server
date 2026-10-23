package display

import (
	"archive/zip"
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestConfigurationInheritance(t *testing.T) {
	c, err := MergeConfiguration(map[string]any{"idle_seconds": 80}, map[string]any{"clock": "combined"})
	if err != nil || c.IdleSeconds != 80 || c.Clock != "combined" || c.SlideSeconds != 30 || c.Enabled {
		t.Fatalf("%+v %v", c, err)
	}
	c, err = MergeConfiguration(map[string]any{"idle_seconds": 80}, map[string]any{"idle_seconds": nil})
	if err != nil || c.IdleSeconds != 80 {
		t.Fatal(c, err)
	}
	for _, v := range []map[string]any{{"clock": "bad"}, {"idle_seconds": -1}, {"timezone": "not/a/zone"}, {"unknown": 1}, {"shortcuts": []any{map[string]any{"name": "bad", "url": "file:///etc/passwd"}}}} {
		if _, e := MergeConfiguration(v, nil); e == nil {
			t.Fatal("accepted", v)
		}
	}
}
func TestPrivateAtomicCacheAndPIN(t *testing.T) {
	v, err := NewVerifier("123456")
	if err != nil {
		t.Fatal(err)
	}
	if CheckPIN(v, "123456") != nil || CheckPIN(v, "123455") == nil {
		t.Fatal("PIN check")
	}
	if _, err := NewVerifier("12345x"); err == nil {
		t.Fatal("invalid PIN")
	}
	for _, pin := range []string{"1234", "12345678901234567890"} {
		if _, err := NewVerifier(pin); err != nil {
			t.Fatal(err)
		}
	}
	for _, pin := range []string{"123", "123456789012345678901"} {
		if _, err := NewVerifier(pin); err == nil {
			t.Fatal("invalid length accepted")
		}
	}
	p := filepath.Join(t.TempDir(), "cache.json")
	if err := WriteCache(p, []byte(`{"ok":true}`)); err != nil {
		t.Fatal(err)
	}
	st, _ := os.Stat(p)
	if st.Mode().Perm() != 0600 {
		t.Fatal(st.Mode())
	}
}
func TestThemeArchiveBoundary(t *testing.T) {
	pack := func(extra string) []byte {
		var b bytes.Buffer
		w := zip.NewWriter(&b)
		f, _ := w.Create("manifest.json")
		f.Write([]byte(`{"api_version":1,"name":"Starter","entry":"index.html"}`))
		f, _ = w.Create("index.html")
		f.Write([]byte("hello"))
		if extra != "" {
			f, _ = w.Create(extra)
			f.Write([]byte("bad"))
		}
		w.Close()
		return b.Bytes()
	}
	if _, err := ReadTheme(pack("")); err != nil {
		t.Fatal(err)
	}
	for _, p := range []string{"../escape", "/absolute", "a/../../escape", `a\b`, "index.html"} {
		if _, err := ReadTheme(pack(p)); err == nil {
			t.Fatal("accepted", p)
		}
	}
}
