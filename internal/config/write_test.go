package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestErrGlobalConfigMissingMessage(t *testing.T) {
	err := ErrGlobalConfigMissing{Path: "/home/user/.chef/config.json"}
	msg := err.Error()
	if !strings.Contains(msg, "chef config") {
		t.Fatalf("error message = %q, want chef config hint", msg)
	}
	if !strings.Contains(msg, "chef config --project") {
		t.Fatalf("error message = %q, want project hint", msg)
	}
}

func TestWriteFileRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "chef", "config.json")
	cfg := Defaults()
	cfg.Provider = "anthropic"
	cfg.Model = "claude-sonnet-4-20250514"

	if err := WriteFile(path, cfg); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var got Config
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if got.Provider != "anthropic" {
		t.Fatalf("provider = %q, want anthropic", got.Provider)
	}
	if got.Model != "claude-sonnet-4-20250514" {
		t.Fatalf("model = %q", got.Model)
	}
	if _, err := os.Stat(path + ".tmp"); !os.IsNotExist(err) {
		t.Fatal("temp file should not remain after write")
	}
}

func TestWriteFileCreatesParentDir(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "nested", "config.json")
	if err := WriteFile(path, Defaults()); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}
}

func TestGlobalConfigPath(t *testing.T) {
	root := t.TempDir()
	t.Setenv("CHEF_DIR", root)
	path, err := GlobalConfigPath()
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(root, "config.json")
	if path != want {
		t.Fatalf("path = %q, want %q", path, want)
	}
}

func TestProjectConfigPath(t *testing.T) {
	got := ProjectConfigPath("/repo/root")
	want := filepath.Join("/repo/root", ".chef", "config.json")
	if got != want {
		t.Fatalf("path = %q, want %q", got, want)
	}
}

func TestLoadRequireGlobalMissing(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("CHEF_DIR", filepath.Join(dir, "no-global"))

	_, err := Load(LoadOptions{WorkDir: dir, RequireGlobal: true})
	if err == nil {
		t.Fatal("expected error")
	}
	var missing ErrGlobalConfigMissing
	if !errors.As(err, &missing) {
		t.Fatalf("error = %T: %v, want ErrGlobalConfigMissing", err, err)
	}
}

func TestLoadRequireGlobalPresent(t *testing.T) {
	dir := t.TempDir()
	globalDir := filepath.Join(dir, "global")
	t.Setenv("CHEF_DIR", globalDir)
	if err := os.MkdirAll(globalDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(globalDir, "config.json"), []byte(`{"model":"custom"}`), 0o644); err != nil {
		t.Fatal(err)
	}

	result, err := Load(LoadOptions{WorkDir: dir, RequireGlobal: true})
	if err != nil {
		t.Fatal(err)
	}
	if result.Config.Model != "custom" {
		t.Fatalf("model = %q, want custom", result.Config.Model)
	}
}
