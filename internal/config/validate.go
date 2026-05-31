package config

import (
	"fmt"
	"strings"

	"github.com/mythosmystery/chef/pkg/provider"
)

var knownContextFiles = map[string]struct{}{
	"project.md": {}, "glossary.md": {}, "features.md": {}, "conventions.md": {},
	"architecture.md": {}, "data.md": {}, "api.md": {}, "workflows.md": {},
}

var knownProviders = map[string]struct{}{
	"openai":    {},
	"anthropic": {},
	"custom":    {},
}

const providerList = "openai|anthropic|custom"

var knownTools = map[string]struct{}{
	"read": {}, "write": {}, "edit": {}, "bash": {}, "grep": {},
	"find": {}, "ls": {}, "context": {}, "diff": {},
}

var knownThemes = map[string]struct{}{
	"dark":  {},
	"light": {},
}

// Validate checks cfg and returns an error if any field is invalid.
func Validate(cfg *Config) error {
	if cfg.Provider == "" {
		return fmt.Errorf("config: provider is required")
	}
	if _, ok := knownProviders[cfg.Provider]; !ok {
		return fmt.Errorf("config: invalid provider %q (want %s)", cfg.Provider, providerList)
	}
	if cfg.Model == "" {
		return fmt.Errorf("config: model is required")
	}
	if err := validateProviderSettings(cfg.Provider, cfg.ActiveProviderConfig()); err != nil {
		return err
	}

	if cfg.Light != nil {
		if cfg.Light.Provider != "" {
			if _, ok := knownProviders[cfg.Light.Provider]; !ok {
				return fmt.Errorf("config: invalid light.provider %q (want %s)", cfg.Light.Provider, providerList)
			}
		}
		lp, lm := cfg.LightModel()
		if lp == "" {
			return fmt.Errorf("config: light provider is required")
		}
		if lm == "" {
			return fmt.Errorf("config: light model is required")
		}
		lightName, _ := cfg.LightModel()
		if err := validateProviderSettings(lightName, cfg.LightProviderConfig()); err != nil {
			return err
		}
	}

	if err := validateSampling(cfg.Sampling); err != nil {
		return err
	}
	if cfg.MaxTurns < 1 {
		return fmt.Errorf("config: maxTurns must be >= 1, got %d", cfg.MaxTurns)
	}
	for name, pc := range cfg.Providers {
		if err := validateProviderEntry(name, pc); err != nil {
			return err
		}
	}

	if err := validateThinking(cfg.Thinking); err != nil {
		return err
	}

	for _, tool := range cfg.Tools {
		if _, ok := knownTools[tool]; !ok {
			return fmt.Errorf("config: unknown tool %q", tool)
		}
	}

	if cfg.MaxConcurrentAgents < 1 {
		return fmt.Errorf("config: maxConcurrentAgents must be >= 1, got %d", cfg.MaxConcurrentAgents)
	}
	if cfg.AgentTimeout.Duration <= 0 {
		return fmt.Errorf("config: agentTimeout must be > 0")
	}

	if cfg.Session.CompactThreshold <= 0 || cfg.Session.CompactThreshold > 1 {
		return fmt.Errorf("config: session.compactThreshold must be in (0, 1], got %v", cfg.Session.CompactThreshold)
	}
	if cfg.Session.CompactMaxTurns < 1 {
		return fmt.Errorf("config: session.compactMaxTurns must be >= 1, got %d", cfg.Session.CompactMaxTurns)
	}
	if cfg.Session.Dir == "" {
		return fmt.Errorf("config: session.dir is required")
	}

	if cfg.ContextFiles.Dir == "" {
		return fmt.Errorf("config: contextFiles.dir is required")
	}
	if err := validateBudget(cfg.ContextFiles.Budget); err != nil {
		return err
	}

	if _, ok := knownThemes[cfg.Theme]; !ok {
		return fmt.Errorf("config: invalid theme %q (want dark|light)", cfg.Theme)
	}

	_ = provider.ParseThinking(cfg.Thinking)
	return nil
}

func validateProviderSettings(providerName string, pc ProviderConfig) error {
	if providerName == "custom" && pc.BaseURL == "" {
		return fmt.Errorf("config: providers.custom.baseURL is required when provider is custom")
	}
	return nil
}

func validateProviderEntry(name string, pc ProviderConfig) error {
	if pc.Timeout.Duration < 0 {
		return fmt.Errorf("config: providers.%s.timeout must be >= 0", name)
	}
	if pc.MaxRetries < 0 {
		return fmt.Errorf("config: providers.%s.maxRetries must be >= 0", name)
	}
	return nil
}

func validateSampling(s SamplingConfig) error {
	if s.Temperature != nil {
		t := *s.Temperature
		if t < 0 || t > 2 {
			return fmt.Errorf("config: sampling.temperature must be in [0, 2], got %v", t)
		}
	}
	if s.TopP != nil {
		p := *s.TopP
		if p < 0 || p > 1 {
			return fmt.Errorf("config: sampling.topP must be in [0, 1], got %v", p)
		}
	}
	if s.TopK != nil && *s.TopK < 0 {
		return fmt.Errorf("config: sampling.topK must be >= 0, got %d", *s.TopK)
	}
	if s.MaxTokens != nil && *s.MaxTokens <= 0 {
		return fmt.Errorf("config: sampling.maxTokens must be > 0, got %d", *s.MaxTokens)
	}
	return nil
}

func validateThinking(s string) error {
	switch provider.Thinking(s) {
	case provider.ThinkingOff, provider.ThinkingLow, provider.ThinkingMedium, provider.ThinkingHigh:
		return nil
	default:
		return fmt.Errorf("config: invalid thinking %q (want off|low|medium|high)", s)
	}
}

func validateBudget(budget map[string]int) error {
	for name, limit := range budget {
		if _, ok := knownContextFiles[name]; !ok {
			return fmt.Errorf("config: unknown context file budget key %q", name)
		}
		if limit <= 0 {
			return fmt.Errorf("config: budget for %q must be > 0, got %d", name, limit)
		}
	}
	return nil
}

// ValidateThinkingFlag validates a CLI thinking override.
func ValidateThinkingFlag(s string) error {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	return validateThinking(s)
}
