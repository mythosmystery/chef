// Package config loads and merges global and project JSON configuration.
package config

// Config is the merged chef configuration.
type Config struct {
	Provider            string                   `json:"provider"`
	Model               string                   `json:"model"`
	Providers           map[string]ProviderConfig `json:"providers,omitempty"`
	Light               *LightConfig             `json:"light,omitempty"`
	Thinking            string                   `json:"thinking"`
	Sampling            SamplingConfig           `json:"sampling,omitempty"`
	MaxTurns            int                      `json:"maxTurns"`
	Tools               []string                 `json:"tools"`
	MaxConcurrentAgents int                      `json:"maxConcurrentAgents"`
	AgentTimeout        Duration                 `json:"agentTimeout"`
	ContextFiles        ContextConfig            `json:"contextFiles"`
	Session             SessionConfig            `json:"session"`
	Bash                BashConfig               `json:"bash"`
	Theme               string                   `json:"theme"`
}

// ProviderConfig holds per-provider connection settings.
// Provider-specific fields are ignored by backends that do not use them.
type ProviderConfig struct {
	BaseURL      string            `json:"baseURL,omitempty"`
	APIKeyEnv    string            `json:"apiKeyEnv,omitempty"`
	Organization string            `json:"organization,omitempty"`
	Project      string            `json:"project,omitempty"`
	Version      string            `json:"version,omitempty"`
	Beta         []string          `json:"beta,omitempty"`
	Headers      map[string]string `json:"headers,omitempty"`
	Timeout      Duration          `json:"timeout,omitempty"`
	MaxRetries   int               `json:"maxRetries,omitempty"`
}

// SamplingConfig holds LLM sampling parameters. Pointer fields mean unset (use model default).
type SamplingConfig struct {
	Temperature *float64 `json:"temperature,omitempty"`
	TopP        *float64 `json:"topP,omitempty"`
	TopK        *int     `json:"topK,omitempty"`
	MaxTokens   *int     `json:"maxTokens,omitempty"`
}

// ActiveProviderConfig returns connection settings for the main provider.
func (c Config) ActiveProviderConfig() ProviderConfig {
	return c.providerConfig(c.Provider)
}

// LightProviderConfig returns connection settings for the light (mini-agent) provider.
func (c Config) LightProviderConfig() ProviderConfig {
	name, _ := c.LightModel()
	return c.providerConfig(name)
}

func (c Config) providerConfig(name string) ProviderConfig {
	if c.Providers == nil {
		return ProviderConfig{}
	}
	return c.Providers[name]
}

// LightConfig overrides provider/model for mini-agents.
type LightConfig struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
}

// ContextConfig configures .chef context files.
type ContextConfig struct {
	Dir    string         `json:"dir"`
	Budget map[string]int `json:"budget"`
}

// SessionConfig configures session persistence and compaction.
type SessionConfig struct {
	Dir              string  `json:"dir"`
	AutoCompact      bool    `json:"autoCompact"`
	CompactThreshold float64 `json:"compactThreshold"`
	CompactMaxTurns  int     `json:"compactMaxTurns"`
}

// BashConfig configures shell command safety.
type BashConfig struct {
	Blocklist []string `json:"blocklist"`
}
