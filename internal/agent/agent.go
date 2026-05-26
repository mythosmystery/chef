// Package agent implements the chef agent core: turns, plans, and mini-agents.
package agent

import (
	"context"

	"github.com/mythosmystery/chef/internal/bus"
	"github.com/mythosmystery/chef/internal/config"
	"github.com/mythosmystery/chef/internal/projctx"
	"github.com/mythosmystery/chef/internal/session"
	"github.com/mythosmystery/chef/internal/tool"
	"github.com/mythosmystery/chef/pkg/provider"
)

// Agent is the main coding agent.
type Agent struct {
	cfg      *config.Config
	provider provider.Provider
	tools    *tool.Registry
	sessions *session.Store
	projctx  *projctx.Manager
	bus      *bus.Bus
}

// New creates a main agent.
func New(cfg *config.Config, p provider.Provider, tools *tool.Registry, sessions *session.Store, proj *projctx.Manager, b *bus.Bus) *Agent {
	return &Agent{
		cfg:      cfg,
		provider: p,
		tools:    tools,
		sessions: sessions,
		projctx:  proj,
		bus:      b,
	}
}

// Run starts the agent event loop for a session.
func (a *Agent) Run(ctx context.Context, sess session.Session) error {
	panic("not implemented")
}

// Stop cancels in-flight agent work.
func (a *Agent) Stop() {
	panic("not implemented")
}
