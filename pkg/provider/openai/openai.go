// Package openai implements the OpenAI provider backend.
package openai

import (
	"context"

	"github.com/mythosmystery/chef/pkg/provider"
)

// Provider implements provider.Provider for OpenAI.
type Provider struct {
	apiKey string
}

// New creates an OpenAI provider.
func New(apiKey string) *Provider {
	return &Provider{apiKey: apiKey}
}

// Name returns the provider name.
func (p *Provider) Name() string { return "openai" }

// Complete streams a completion from OpenAI.
func (p *Provider) Complete(ctx context.Context, req provider.Request) (<-chan provider.StreamEvent, error) {
	panic("not implemented")
}
