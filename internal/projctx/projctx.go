// Package projctx manages .chef/*.md project context files.
package projctx

import (
	"github.com/mythosmystery/chef/internal/config"
	"github.com/mythosmystery/chef/internal/tokens"
	"github.com/mythosmystery/chef/internal/workspace"
)

// Manager coordinates context file I/O, budget, and operations.
type Manager struct {
	cfg       config.ContextConfig
	ws        *workspace.Workspace
	counter   tokens.Counter
	store     *Store
}

// NewManager creates a project context manager.
func NewManager(cfg config.ContextConfig, ws *workspace.Workspace, counter tokens.Counter) *Manager {
	return &Manager{
		cfg:     cfg,
		ws:      ws,
		counter: counter,
		store:   NewStore(cfg.Dir),
	}
}
