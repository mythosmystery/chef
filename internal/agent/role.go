package agent

import (
	"time"

	"github.com/mythosmystery/chef/internal/agent/spawn"
)

// RoleKind identifies mini-agent roles.
type RoleKind string

const (
	RoleWork       RoleKind = "work"
	RoleContextMgr RoleKind = "context_manager"
	RoleInit       RoleKind = "init"
)

// RoleConfig defines a mini-agent role's prompt, tools, and defaults.
type RoleConfig struct {
	Kind         RoleKind
	SystemPrompt string
	Tools        []string
	Timeout      time.Duration
	Priority     spawn.Priority
}

// Role returns the configuration for a mini-agent role.
func Role(kind RoleKind) RoleConfig {
	panic("not implemented")
}
