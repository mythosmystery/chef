// Package provider defines the public LLM provider abstraction for chef.
package provider

import "context"

// Provider streams completions from an LLM backend.
type Provider interface {
	Name() string
	Complete(ctx context.Context, req Request) (<-chan StreamEvent, error)
}

// Request is a completion request.
type Request struct {
	Model    string
	Messages []Message
	Tools    []ToolDef
	Thinking Thinking
}

// LightProvider resolves the light model configuration for mini-agents.
type LightProvider interface {
	LightModel(main Provider) (Provider, error)
}
