package config

import (
	"testing"
	"time"
)

func TestMergeLightPartial(t *testing.T) {
	base := Defaults()
	base.Light = &LightConfig{Provider: "openai", Model: "gpt-4o-mini"}

	t.Run("model only", func(t *testing.T) {
		got := Merge(base, fileConfig{
			Light: &lightOverlay{Model: strPtr("gpt-3.5-turbo")},
		})
		if got.Light.Provider != "openai" {
			t.Fatalf("provider = %q, want openai", got.Light.Provider)
		}
		if got.Light.Model != "gpt-3.5-turbo" {
			t.Fatalf("model = %q, want gpt-3.5-turbo", got.Light.Model)
		}
	})

	t.Run("provider only", func(t *testing.T) {
		got := Merge(base, fileConfig{
			Light: &lightOverlay{Provider: strPtr("anthropic")},
		})
		if got.Light.Provider != "anthropic" {
			t.Fatalf("provider = %q, want anthropic", got.Light.Provider)
		}
		if got.Light.Model != "gpt-4o-mini" {
			t.Fatalf("model = %q, want gpt-4o-mini", got.Light.Model)
		}
	})

	t.Run("both", func(t *testing.T) {
		got := Merge(base, fileConfig{
			Light: &lightOverlay{
				Provider: strPtr("anthropic"),
				Model:    strPtr("claude-haiku"),
			},
		})
		if got.Light.Provider != "anthropic" || got.Light.Model != "claude-haiku" {
			t.Fatalf("light = %+v", got.Light)
		}
	})
}

func TestMergeToolsEmptyVsAbsent(t *testing.T) {
	base := Defaults()

	absent := Merge(base, fileConfig{})
	if len(absent.Tools) != len(base.Tools) {
		t.Fatalf("absent tools changed: got %d want %d", len(absent.Tools), len(base.Tools))
	}

	empty := Merge(base, fileConfig{Tools: &[]string{}})
	if len(empty.Tools) != 0 {
		t.Fatalf("empty tools = %v, want []", empty.Tools)
	}
}

func TestMergeBudgetMap(t *testing.T) {
	base := Defaults()
	got := Merge(base, fileConfig{
		ContextFiles: &contextOverlay{
			Budget: map[string]int{"project.md": 3000},
		},
	})

	if got.ContextFiles.Budget["project.md"] != 3000 {
		t.Fatalf("project.md budget = %d, want 3000", got.ContextFiles.Budget["project.md"])
	}
	if got.ContextFiles.Budget["glossary.md"] != 1000 {
		t.Fatalf("glossary.md budget = %d, want 1000 (preserved from base)", got.ContextFiles.Budget["glossary.md"])
	}
}

func TestMergeAutoCompactFalse(t *testing.T) {
	base := Defaults()
	if !base.Session.AutoCompact {
		t.Fatal("base autoCompact should be true")
	}

	got := Merge(base, fileConfig{
		Session: &sessionOverlay{AutoCompact: boolPtr(false)},
	})
	if got.Session.AutoCompact {
		t.Fatal("autoCompact should be false after override")
	}
}

func TestMergeAgentTimeout(t *testing.T) {
	base := Defaults()
	got := Merge(base, fileConfig{
		AgentTimeout: &Duration{Duration: 10 * time.Minute},
	})
	if got.AgentTimeout.Duration != 10*time.Minute {
		t.Fatalf("agentTimeout = %v, want 10m", got.AgentTimeout.Duration)
	}
}

func TestLightModelFallback(t *testing.T) {
	base := Defaults()
	p, m := base.LightModel()
	if p != "openai" || m != "gpt-4o" {
		t.Fatalf("no light: got (%q, %q)", p, m)
	}

	withLight := base
	withLight.Light = &LightConfig{Model: "mini"}
	p, m = withLight.LightModel()
	if p != "openai" || m != "mini" {
		t.Fatalf("partial light: got (%q, %q)", p, m)
	}
}

func strPtr(s string) *string { return &s }
func boolPtr(b bool) *bool    { return &b }
