package cli

import (
	"flag"
	"testing"

	"github.com/mythosmystery/chef/internal/config"
)

func TestBuildConfigFromAnswers(t *testing.T) {
	base := config.Defaults()
	cfg := buildConfigFromAnswers(base, configAnswers{
		Provider: "anthropic",
		Model:    "claude-sonnet-4-20250514",
		Thinking: "high",
		Theme:    "light",
	})

	if cfg.Provider != "anthropic" {
		t.Fatalf("provider = %q", cfg.Provider)
	}
	if cfg.Model != "claude-sonnet-4-20250514" {
		t.Fatalf("model = %q", cfg.Model)
	}
	if cfg.Thinking != "high" {
		t.Fatalf("thinking = %q", cfg.Thinking)
	}
	if cfg.Theme != "light" {
		t.Fatalf("theme = %q", cfg.Theme)
	}
	if cfg.Light != nil {
		t.Fatal("light should be nil when disabled")
	}
}

func TestBuildConfigFromAnswersWithLight(t *testing.T) {
	base := config.Defaults()
	cfg := buildConfigFromAnswers(base, configAnswers{
		Provider:      "openai",
		Model:         "gpt-4o",
		Thinking:      "medium",
		Theme:         "dark",
		LightEnabled:  true,
		LightProvider: "openai",
		LightModel:    "gpt-4o-mini",
	})

	if cfg.Light == nil {
		t.Fatal("light should be set")
	}
	if cfg.Light.Model != "gpt-4o-mini" {
		t.Fatalf("light model = %q", cfg.Light.Model)
	}
}

func TestParseConfigFlags(t *testing.T) {
	fs := flag.NewFlagSet("config", flag.ContinueOnError)
	project := fs.Bool("project", false, "")
	help := fs.Bool("h", false, "")
	fs.BoolVar(help, "help", false, "")

	if err := fs.Parse([]string{"--project"}); err != nil {
		t.Fatal(err)
	}
	if !*project {
		t.Fatal("project should be true")
	}
	if *help {
		t.Fatal("help should be false")
	}
}

func TestAnswersFromConfig(t *testing.T) {
	base := config.Defaults()
	base.Light = &config.LightConfig{Provider: "openai", Model: "gpt-4o-mini"}

	a := answersFromConfig(base)
	if !a.LightEnabled {
		t.Fatal("light should be enabled")
	}
	if a.LightModel != "gpt-4o-mini" {
		t.Fatalf("light model = %q", a.LightModel)
	}
}

func TestPrintAPIKeyReminderProviders(t *testing.T) {
	t.Run("openai only", func(t *testing.T) {
		cfg := config.Defaults()
		printAPIKeyReminder(&cfg)
	})
	t.Run("anthropic only", func(t *testing.T) {
		cfg := config.Defaults()
		cfg.Provider = "anthropic"
		printAPIKeyReminder(&cfg)
	})
	t.Run("mixed light provider", func(t *testing.T) {
		cfg := config.Defaults()
		cfg.Light = &config.LightConfig{Provider: "anthropic", Model: "claude-haiku"}
		printAPIKeyReminder(&cfg)
	})
	t.Run("custom provider", func(t *testing.T) {
		cfg := config.Defaults()
		cfg.Provider = "custom"
		cfg.Providers["custom"] = config.ProviderConfig{
			BaseURL:   "https://example.com/v1",
			APIKeyEnv: "MY_KEY",
		}
		printAPIKeyReminder(&cfg)
	})
}

func TestBuildConfigFromAnswersCustom(t *testing.T) {
	base := config.Defaults()
	cfg := buildConfigFromAnswers(base, configAnswers{
		Provider:        "custom",
		CustomBaseURL:   "https://example.com/v1",
		CustomAPIKeyEnv: "MY_KEY",
		Model:           "llama-3.1-70b",
		Thinking:        "medium",
		Theme:           "dark",
	})
	if cfg.Providers["custom"].BaseURL != "https://example.com/v1" {
		t.Fatalf("custom baseURL = %q", cfg.Providers["custom"].BaseURL)
	}
	if err := config.Validate(&cfg); err != nil {
		t.Fatalf("validate: %v", err)
	}
}
