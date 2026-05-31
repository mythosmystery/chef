package config

import "github.com/mythosmystery/chef/pkg/provider"

// Merge overlays fileConfig onto base and returns the merged config.
func Merge(base Config, overlay fileConfig) Config {
	out := base

	if overlay.Provider != nil {
		out.Provider = *overlay.Provider
	}
	if overlay.Model != nil {
		out.Model = *overlay.Model
	}
	if len(overlay.Providers) > 0 {
		out.Providers = mergeProviders(out.Providers, overlay.Providers)
	}
	if overlay.Sampling != nil {
		out.Sampling = mergeSampling(out.Sampling, *overlay.Sampling)
	}
	if overlay.MaxTurns != nil {
		out.MaxTurns = *overlay.MaxTurns
	}
	if overlay.Thinking != nil {
		out.Thinking = *overlay.Thinking
	}
	if overlay.Tools != nil {
		out.Tools = append([]string(nil), (*overlay.Tools)...)
	}
	if overlay.MaxConcurrentAgents != nil {
		out.MaxConcurrentAgents = *overlay.MaxConcurrentAgents
	}
	if overlay.AgentTimeout != nil {
		out.AgentTimeout = *overlay.AgentTimeout
	}
	if overlay.Theme != nil {
		out.Theme = *overlay.Theme
	}

	out.Light = mergeLight(out.Light, overlay.Light)

	if overlay.ContextFiles != nil {
		out.ContextFiles = mergeContextFiles(out.ContextFiles, *overlay.ContextFiles)
	}
	if overlay.Session != nil {
		out.Session = mergeSession(out.Session, *overlay.Session)
	}
	if overlay.Bash != nil {
		out.Bash = mergeBash(out.Bash, *overlay.Bash)
	}

	return out
}

func mergeProviders(base map[string]ProviderConfig, overlay map[string]providerOverlay) map[string]ProviderConfig {
	if base == nil {
		base = make(map[string]ProviderConfig, len(overlay))
	}
	out := make(map[string]ProviderConfig, len(base)+len(overlay))
	for k, v := range base {
		out[k] = v
	}
	for name, o := range overlay {
		out[name] = mergeProvider(out[name], o)
	}
	return out
}

func mergeProvider(base ProviderConfig, overlay providerOverlay) ProviderConfig {
	out := base
	if overlay.BaseURL != nil {
		out.BaseURL = *overlay.BaseURL
	}
	if overlay.APIKeyEnv != nil {
		out.APIKeyEnv = *overlay.APIKeyEnv
	}
	if overlay.Organization != nil {
		out.Organization = *overlay.Organization
	}
	if overlay.Project != nil {
		out.Project = *overlay.Project
	}
	if overlay.Version != nil {
		out.Version = *overlay.Version
	}
	if overlay.Beta != nil {
		out.Beta = append([]string(nil), (*overlay.Beta)...)
	}
	if len(overlay.Headers) > 0 {
		if out.Headers == nil {
			out.Headers = make(map[string]string, len(overlay.Headers))
		}
		for k, v := range overlay.Headers {
			out.Headers[k] = v
		}
	}
	if overlay.Timeout != nil {
		out.Timeout = *overlay.Timeout
	}
	if overlay.MaxRetries != nil {
		out.MaxRetries = *overlay.MaxRetries
	}
	return out
}

func mergeSampling(base SamplingConfig, overlay samplingOverlay) SamplingConfig {
	out := base
	if overlay.Temperature != nil {
		out.Temperature = overlay.Temperature
	}
	if overlay.TopP != nil {
		out.TopP = overlay.TopP
	}
	if overlay.TopK != nil {
		out.TopK = overlay.TopK
	}
	if overlay.MaxTokens != nil {
		out.MaxTokens = overlay.MaxTokens
	}
	return out
}

func mergeLight(base *LightConfig, overlay *lightOverlay) *LightConfig {
	if overlay == nil {
		return base
	}

	var out LightConfig
	if base != nil {
		out = *base
	}
	if overlay.Provider != nil {
		out.Provider = *overlay.Provider
	}
	if overlay.Model != nil {
		out.Model = *overlay.Model
	}
	if out.Provider == "" && out.Model == "" {
		return nil
	}
	return &out
}

func mergeContextFiles(base ContextConfig, overlay contextOverlay) ContextConfig {
	out := base
	if overlay.Dir != nil {
		out.Dir = *overlay.Dir
	}
	if len(overlay.Budget) > 0 {
		if out.Budget == nil {
			out.Budget = make(map[string]int, len(overlay.Budget))
		}
		for k, v := range overlay.Budget {
			out.Budget[k] = v
		}
	}
	return out
}

func mergeSession(base SessionConfig, overlay sessionOverlay) SessionConfig {
	out := base
	if overlay.Dir != nil {
		out.Dir = *overlay.Dir
	}
	if overlay.AutoCompact != nil {
		out.AutoCompact = *overlay.AutoCompact
	}
	if overlay.CompactThreshold != nil {
		out.CompactThreshold = *overlay.CompactThreshold
	}
	if overlay.CompactMaxTurns != nil {
		out.CompactMaxTurns = *overlay.CompactMaxTurns
	}
	return out
}

func mergeBash(base BashConfig, overlay bashOverlay) BashConfig {
	out := base
	if overlay.Blocklist != nil {
		out.Blocklist = append([]string(nil), (*overlay.Blocklist)...)
	}
	return out
}

// MainModel returns the main agent provider and model.
func (c Config) MainModel() (providerName, model string) {
	return c.Provider, c.Model
}

// LightModel returns the mini-agent provider and model, falling back to main.
func (c Config) LightModel() (providerName, model string) {
	if c.Light == nil {
		return c.Provider, c.Model
	}
	providerName = c.Light.Provider
	if providerName == "" {
		providerName = c.Provider
	}
	model = c.Light.Model
	if model == "" {
		model = c.Model
	}
	return providerName, model
}

// ThinkingLevel parses the configured thinking level.
func (c Config) ThinkingLevel() provider.Thinking {
	return provider.ParseThinking(c.Thinking)
}
