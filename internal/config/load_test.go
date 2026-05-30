package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestLoadDefaultsOnly(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("CHEF_DIR", filepath.Join(dir, "no-global"))

	result, err := Load(LoadOptions{WorkDir: dir})
	if err != nil {
		t.Fatal(err)
	}
	cfg := result.Config
	if cfg.Provider != "openai" {
		t.Fatalf("provider = %q, want openai", cfg.Provider)
	}
	if cfg.MaxConcurrentAgents != 4 {
		t.Fatalf("maxConcurrentAgents = %d, want 4", cfg.MaxConcurrentAgents)
	}
}

func TestLoadGlobalAndProjectMerge(t *testing.T) {
	root := t.TempDir()
	globalDir := filepath.Join(root, "global")
	projectRoot := filepath.Join(root, "repo")
	subDir := filepath.Join(projectRoot, "internal", "pkg")
	if err := os.MkdirAll(subDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(projectRoot, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	chefDir := filepath.Join(projectRoot, ".chef")
	if err := os.MkdirAll(chefDir, 0o755); err != nil {
		t.Fatal(err)
	}

	t.Setenv("CHEF_DIR", globalDir)
	if err := os.MkdirAll(globalDir, 0o755); err != nil {
		t.Fatal(err)
	}
	globalJSON := `{"thinking":"high","session":{"autoCompact":true}}`
	if err := os.WriteFile(filepath.Join(globalDir, "config.json"), []byte(globalJSON), 0o644); err != nil {
		t.Fatal(err)
	}

	projectJSON := `{
		"model": "gpt-4o-mini",
		"session": {"autoCompact": false},
		"contextFiles": {"budget": {"project.md": 2500}}
	}`
	if err := os.WriteFile(filepath.Join(chefDir, "config.json"), []byte(projectJSON), 0o644); err != nil {
		t.Fatal(err)
	}

	result, err := Load(LoadOptions{WorkDir: subDir})
	if err != nil {
		t.Fatal(err)
	}
	cfg := result.Config

	if cfg.Model != "gpt-4o-mini" {
		t.Fatalf("model = %q, want gpt-4o-mini", cfg.Model)
	}
	if cfg.Thinking != "high" {
		t.Fatalf("thinking = %q, want high (from global)", cfg.Thinking)
	}
	if cfg.Session.AutoCompact {
		t.Fatal("autoCompact should be false from project override")
	}
	if cfg.ContextFiles.Budget["project.md"] != 2500 {
		t.Fatalf("project.md budget = %d, want 2500", cfg.ContextFiles.Budget["project.md"])
	}
	if cfg.ContextFiles.Budget["glossary.md"] != 1000 {
		t.Fatalf("glossary.md budget = %d, want 1000 (preserved)", cfg.ContextFiles.Budget["glossary.md"])
	}
	if result.ProjectRoot != projectRoot {
		t.Fatalf("project root = %q, want %q", result.ProjectRoot, projectRoot)
	}
}

func TestLoadAgentTimeoutDurationString(t *testing.T) {
	root := t.TempDir()
	t.Setenv("CHEF_DIR", filepath.Join(root, "global"))
	projectRoot := filepath.Join(root, "repo")
	if err := os.MkdirAll(filepath.Join(projectRoot, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	chefDir := filepath.Join(projectRoot, ".chef")
	if err := os.MkdirAll(chefDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(chefDir, "config.json"), []byte(`{"agentTimeout":"10m"}`), 0o644); err != nil {
		t.Fatal(err)
	}

	result, err := Load(LoadOptions{WorkDir: projectRoot})
	if err != nil {
		t.Fatal(err)
	}
	if result.Config.AgentTimeout.Duration != 10*time.Minute {
		t.Fatalf("agentTimeout = %v, want 10m", result.Config.AgentTimeout.Duration)
	}
}

func TestLoadSessionDirTildeExpansion(t *testing.T) {
	root := t.TempDir()
	t.Setenv("CHEF_DIR", filepath.Join(root, "global"))
	projectRoot := filepath.Join(root, "repo")
	if err := os.MkdirAll(filepath.Join(projectRoot, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}

	result, err := Load(LoadOptions{WorkDir: projectRoot})
	if err != nil {
		t.Fatal(err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(home, ".chef", "sessions")
	if result.Config.Session.Dir != want {
		t.Fatalf("session.dir = %q, want %q", result.Config.Session.Dir, want)
	}
}

func TestLoadInvalidThinking(t *testing.T) {
	root := t.TempDir()
	globalDir := filepath.Join(root, "global")
	t.Setenv("CHEF_DIR", globalDir)
	if err := os.MkdirAll(globalDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(globalDir, "config.json"), []byte(`{"thinking":"banana"}`), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(LoadOptions{WorkDir: root})
	if err == nil {
		t.Fatal("expected error for invalid thinking")
	}
	if !strings.Contains(err.Error(), "invalid thinking") {
		t.Fatalf("error = %v", err)
	}
}

func TestLoadCLIModelOverride(t *testing.T) {
	root := t.TempDir()
	projectRoot := filepath.Join(root, "repo")
	chefDir := filepath.Join(projectRoot, ".chef")
	if err := os.MkdirAll(filepath.Join(projectRoot, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(chefDir, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("CHEF_DIR", filepath.Join(root, "global"))
	if err := os.WriteFile(filepath.Join(chefDir, "config.json"), []byte(`{"model":"from-file"}`), 0o644); err != nil {
		t.Fatal(err)
	}

	result, err := Load(LoadOptions{
		WorkDir: projectRoot,
		Flags:   FlagOverrides{Model: "from-cli"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Config.Model != "from-cli" {
		t.Fatalf("model = %q, want from-cli", result.Config.Model)
	}
}

func TestLoadNoContextFlag(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("CHEF_DIR", filepath.Join(dir, "global"))

	result, err := Load(LoadOptions{
		WorkDir: dir,
		Flags:   FlagOverrides{NoContext: true},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !result.NoContext {
		t.Fatal("NoContext should be true")
	}
}

func TestLoadMalformedJSON(t *testing.T) {
	root := t.TempDir()
	globalDir := filepath.Join(root, "global")
	t.Setenv("CHEF_DIR", globalDir)
	if err := os.MkdirAll(globalDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(globalDir, "config.json"), []byte(`{bad`), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(LoadOptions{WorkDir: root})
	if err == nil {
		t.Fatal("expected JSON error")
	}
}

func TestApplyFlagsNoTools(t *testing.T) {
	cfg := Defaults()
	if err := ApplyFlags(&cfg, FlagOverrides{NoTools: true}); err != nil {
		t.Fatal(err)
	}
	if len(cfg.Tools) != 0 {
		t.Fatalf("tools = %v, want empty", cfg.Tools)
	}
}
