package app

import (
	"os"
	"testing"

	"github.com/mythosmystery/chef/internal/config"
	"github.com/mythosmystery/chef/pkg/provider/anthropic"
	"github.com/mythosmystery/chef/pkg/provider/openai"
)

func TestBuildProviderOpenAI(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "sk-test")
	cfg := config.Defaults()
	p, err := buildProvider("openai", cfg.ActiveProviderConfig())
	if err != nil {
		t.Fatal(err)
	}
	if p.Name() != "openai" {
		t.Fatalf("name = %q, want openai", p.Name())
	}
}

func TestBuildProviderCustom(t *testing.T) {
	t.Setenv("MY_KEY", "secret")
	cfg := config.Defaults()
	cfg.Provider = "custom"
	cfg.Providers["custom"] = config.ProviderConfig{
		BaseURL:   "https://example.com/v1",
		APIKeyEnv: "MY_KEY",
	}
	p, err := buildProvider("custom", cfg.Providers["custom"])
	if err != nil {
		t.Fatal(err)
	}
	if p.Name() != "custom" {
		t.Fatalf("name = %q, want custom", p.Name())
	}
	o, ok := p.(*openai.Provider)
	if !ok {
		t.Fatalf("type = %T, want *openai.Provider", p)
	}
	if o.BaseURL() != "https://example.com/v1" {
		t.Fatalf("baseURL = %q", o.BaseURL())
	}
}

func TestBuildProviderAnthropic(t *testing.T) {
	t.Setenv("ANTHROPIC_API_KEY", "sk-ant-test")
	cfg := config.Defaults()
	cfg.Provider = "anthropic"
	p, err := buildProvider("anthropic", cfg.ActiveProviderConfig())
	if err != nil {
		t.Fatal(err)
	}
	a, ok := p.(*anthropic.Provider)
	if !ok {
		t.Fatalf("type = %T, want *anthropic.Provider", p)
	}
	if a.BaseURL() != "https://api.anthropic.com" {
		t.Fatalf("baseURL = %q", a.BaseURL())
	}
	if a.APIVersion() != "2023-06-01" {
		t.Fatalf("version = %q", a.APIVersion())
	}
}

func TestBuildProviderMissingEnvKey(t *testing.T) {
	_ = os.Unsetenv("MISSING_CHEF_KEY")
	cfg := config.Defaults()
	cfg.Providers["custom"] = config.ProviderConfig{
		BaseURL:   "https://example.com/v1",
		APIKeyEnv: "MISSING_CHEF_KEY",
	}
	_, err := buildProvider("custom", cfg.Providers["custom"])
	if err == nil {
		t.Fatal("expected error for missing env key")
	}
}
