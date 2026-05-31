package config

// fileConfig is used only when reading JSON config files. Pointer fields
// distinguish "absent" from zero values so merge can apply partial overrides.
type fileConfig struct {
	Provider            *string                    `json:"provider"`
	Model               *string                    `json:"model"`
	Providers           map[string]providerOverlay `json:"providers"`
	Light               *lightOverlay              `json:"light"`
	Thinking            *string                    `json:"thinking"`
	Sampling            *samplingOverlay           `json:"sampling"`
	MaxTurns            *int                       `json:"maxTurns"`
	Tools               *[]string                  `json:"tools"`
	MaxConcurrentAgents *int                       `json:"maxConcurrentAgents"`
	AgentTimeout        *Duration                  `json:"agentTimeout"`
	ContextFiles        *contextOverlay            `json:"contextFiles"`
	Session             *sessionOverlay            `json:"session"`
	Bash                *bashOverlay               `json:"bash"`
	Theme               *string                    `json:"theme"`
}

type providerOverlay struct {
	BaseURL      *string           `json:"baseURL"`
	APIKeyEnv    *string           `json:"apiKeyEnv"`
	Organization *string           `json:"organization"`
	Project      *string           `json:"project"`
	Version      *string           `json:"version"`
	Beta         *[]string         `json:"beta"`
	Headers      map[string]string `json:"headers"`
	Timeout      *Duration         `json:"timeout"`
	MaxRetries   *int              `json:"maxRetries"`
}

type samplingOverlay struct {
	Temperature *float64 `json:"temperature"`
	TopP        *float64 `json:"topP"`
	TopK        *int     `json:"topK"`
	MaxTokens   *int     `json:"maxTokens"`
}

type lightOverlay struct {
	Provider *string `json:"provider"`
	Model    *string `json:"model"`
}

type contextOverlay struct {
	Dir    *string        `json:"dir"`
	Budget map[string]int `json:"budget"`
}

type sessionOverlay struct {
	Dir              *string  `json:"dir"`
	AutoCompact      *bool    `json:"autoCompact"`
	CompactThreshold *float64 `json:"compactThreshold"`
	CompactMaxTurns  *int     `json:"compactMaxTurns"`
}

type bashOverlay struct {
	Blocklist *[]string `json:"blocklist"`
}
