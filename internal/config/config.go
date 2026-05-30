// Package config loads and merges global and project JSON configuration.
package config

// Config is the merged chef configuration.
type Config struct {
	Provider            string        `json:"provider"`
	Model               string        `json:"model"`
	Light               *LightConfig  `json:"light,omitempty"`
	Thinking            string        `json:"thinking"`
	Tools               []string      `json:"tools"`
	MaxConcurrentAgents int           `json:"maxConcurrentAgents"`
	AgentTimeout        Duration      `json:"agentTimeout"`
	ContextFiles        ContextConfig `json:"contextFiles"`
	Session             SessionConfig `json:"session"`
	Bash                BashConfig    `json:"bash"`
	Theme               string        `json:"theme"`
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
