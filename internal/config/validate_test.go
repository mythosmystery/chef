package config

import "testing"

func TestValidateCustomRequiresBaseURL(t *testing.T) {
	cfg := Defaults()
	cfg.Provider = "custom"
	cfg.Model = "llama-3.1-70b"
	cfg.Providers["custom"] = ProviderConfig{APIKeyEnv: "MY_KEY"}

	err := Validate(&cfg)
	if err == nil {
		t.Fatal("expected error for custom without baseURL")
	}
}

func TestValidateCustomWithBaseURL(t *testing.T) {
	cfg := Defaults()
	cfg.Provider = "custom"
	cfg.Model = "llama-3.1-70b"
	cfg.Providers["custom"] = ProviderConfig{
		BaseURL:   "https://example.com/v1",
		APIKeyEnv: "MY_KEY",
	}

	if err := Validate(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateSampling(t *testing.T) {
	cfg := Defaults()
	bad := 3.0
	cfg.Sampling.Temperature = &bad
	if err := Validate(&cfg); err == nil {
		t.Fatal("expected temperature validation error")
	}
}

func TestValidateMaxTurns(t *testing.T) {
	cfg := Defaults()
	cfg.MaxTurns = 0
	if err := Validate(&cfg); err == nil {
		t.Fatal("expected maxTurns validation error")
	}
}
