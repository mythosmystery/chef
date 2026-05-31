// Package anthropic implements the Anthropic provider backend.
package anthropic

import (
	"context"
	"strings"
	"time"

	"github.com/mythosmystery/chef/pkg/provider"
)

const (
	defaultBaseURL    = "https://api.anthropic.com"
	defaultAPIVersion = "2023-06-01"
	// DefaultMaxTokens is used when Request.MaxTokens is zero (Anthropic requires max_tokens).
	DefaultMaxTokens = 4096
)

// Options configures an Anthropic provider.
type Options struct {
	APIKey     string
	BaseURL    string
	APIVersion string
	Beta       []string
	Headers    map[string]string
	Timeout    time.Duration
	MaxRetries int
}

// Provider implements provider.Provider for Anthropic.
type Provider struct {
	apiKey     string
	baseURL    string
	apiVersion string
	beta       []string
	headers    map[string]string
	timeout    time.Duration
	maxRetries int
}

// New creates an Anthropic provider.
func New(opts Options) *Provider {
	baseURL := opts.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	apiVersion := opts.APIVersion
	if apiVersion == "" {
		apiVersion = defaultAPIVersion
	}
	beta := opts.Beta
	if beta != nil {
		beta = append([]string(nil), beta...)
	}
	headers := opts.Headers
	if headers == nil {
		headers = make(map[string]string)
	} else {
		headers = copyHeaders(headers)
	}
	return &Provider{
		apiKey:     opts.APIKey,
		baseURL:    baseURL,
		apiVersion: apiVersion,
		beta:       beta,
		headers:    headers,
		timeout:    opts.Timeout,
		maxRetries: opts.MaxRetries,
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
func (p *Provider) Name() string { return "anthropic" }

// BaseURL returns the configured API base URL.
func (p *Provider) BaseURL() string { return p.baseURL }

// APIVersion returns the anthropic-version header value.
func (p *Provider) APIVersion() string { return p.apiVersion }

// BetaHeader returns the anthropic-beta header value (comma-separated).
func (p *Provider) BetaHeader() string {
	if len(p.beta) == 0 {
		return ""
	}
	return strings.Join(p.beta, ",")
}

// EffectiveMaxTokens returns max_tokens for the request, applying DefaultMaxTokens when unset.
func EffectiveMaxTokens(req provider.Request) int {
	if req.MaxTokens > 0 {
		return req.MaxTokens
	}
	return DefaultMaxTokens
}

// Complete streams a completion from Anthropic.
func (p *Provider) Complete(ctx context.Context, req provider.Request) (<-chan provider.StreamEvent, error) {
	panic("not implemented")
}
