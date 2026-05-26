package config

// Merge overlays project overrides onto base.
func Merge(base, override Config) Config {
	out := base
	if override.Provider != "" {
		out.Provider = override.Provider
	}
	if override.Model != "" {
		out.Model = override.Model
	}
	if override.Light != nil {
		out.Light = override.Light
	}
	if override.Thinking != "" {
		out.Thinking = override.Thinking
	}
	if len(override.Tools) > 0 {
		out.Tools = override.Tools
	}
	if override.MaxConcurrentAgents > 0 {
		out.MaxConcurrentAgents = override.MaxConcurrentAgents
	}
	if override.AgentTimeout > 0 {
		out.AgentTimeout = override.AgentTimeout
	}
	if override.Theme != "" {
		out.Theme = override.Theme
	}
	return out
}
