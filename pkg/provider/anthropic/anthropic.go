// Package anthropic implements the Anthropic provider backend.
package anthropic

import (
	"context"

	"github.com/mythosmystery/chef/pkg/provider"
)

// Provider implements provider.Provider for Anthropic.
type Provider struct {
	apiKey string
}

// New creates an Anthropic provider.
func New(apiKey string) *Provider {
	return &Provider{apiKey: apiKey}
}

// Name returns the provider name.
func (p *Provider) Name() string { return "anthropic" }

// Complete streams a completion from Anthropic.
func (p *Provider) Complete(ctx context.Context, req provider.Request) (<-chan provider.StreamEvent, error) {
	panic("not implemented")
}
