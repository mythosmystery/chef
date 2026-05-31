// Package openai implements the OpenAI provider backend.
package openai

import (
	"context"
	"time"

	"github.com/mythosmystery/chef/pkg/provider"
)

// Options configures an OpenAI-compatible provider.
type Options struct {
	APIKey       string
	BaseURL      string
	Organization string
	Project      string
	Headers      map[string]string
	Timeout      time.Duration
	MaxRetries   int
	Name         string // provider name reported by Name(); defaults to "openai"
}

// Provider implements provider.Provider for OpenAI and OpenAI-compatible APIs.
type Provider struct {
	apiKey       string
	baseURL      string
	organization string
	project      string
	headers      map[string]string
	timeout      time.Duration
	maxRetries   int
	name         string
}

// New creates an OpenAI-compatible provider.
func New(opts Options) *Provider {
	name := opts.Name
	if name == "" {
		name = "openai"
	}
	headers := opts.Headers
	if headers == nil {
		headers = make(map[string]string)
	} else {
		headers = copyHeaders(headers)
	}
	return &Provider{
		apiKey:       opts.APIKey,
		baseURL:      opts.BaseURL,
		organization: opts.Organization,
		project:      opts.Project,
		headers:      headers,
		timeout:      opts.Timeout,
		maxRetries:   opts.MaxRetries,
		name:         name,
	}
}

func copyHeaders(h map[string]string) map[string]string {
	out := make(map[string]string, len(h))
	for k, v := range h {
		out[k] = v
	}
	return out
}

// Name returns the provider name.
func (p *Provider) Name() string { return p.name }

// BaseURL returns the configured API base URL.
func (p *Provider) BaseURL() string { return p.baseURL }

// Complete streams a completion from OpenAI.
func (p *Provider) Complete(ctx context.Context, req provider.Request) (<-chan provider.StreamEvent, error) {
	panic("not implemented")
}
