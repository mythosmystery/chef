package config

import "time"

// Defaults returns the default chef configuration.
func Defaults() Config {
	return Config{
		Provider: "openai",
		Model:    "gpt-4o",
		Providers: map[string]ProviderConfig{
			"openai":    {APIKeyEnv: "OPENAI_API_KEY"},
			"anthropic": {
				APIKeyEnv: "ANTHROPIC_API_KEY",
				BaseURL:   "https://api.anthropic.com",
				Version:   "2023-06-01",
			},
		},
		MaxTurns: 50,
		Thinking: "medium",
		Tools:               []string{"read", "write", "edit", "bash", "grep", "find", "ls", "context", "diff"},
		MaxConcurrentAgents: 4,
		AgentTimeout:        Duration{Duration: 5 * time.Minute},
		ContextFiles: ContextConfig{
			Dir: ".chef",
			Budget: map[string]int{
				"project.md":      2000,
				"glossary.md":     1000,
				"features.md":     1500,
				"conventions.md":  1000,
				"architecture.md": 1500,
				"data.md":         1500,
				"api.md":          1500,
				"workflows.md":    500,
			},
		},
		Session: SessionConfig{
			Dir:              "~/.chef/sessions",
			AutoCompact:      true,
			CompactThreshold: 0.8,
			CompactMaxTurns:  50,
		},
		Bash: BashConfig{
			Blocklist: []string{"rm -rf /", "mkfs", "dd if=", "shutdown", "reboot"},
		},
		Theme: "dark",
	}
}
