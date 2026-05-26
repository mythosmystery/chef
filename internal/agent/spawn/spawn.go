// Package spawn implements mini-agent spawning with concurrency and priority.
package spawn

import (
	"context"
	"time"

	"github.com/mythosmystery/chef/internal/tool"
	"github.com/mythosmystery/chef/pkg/provider"
)

// Priority identifies task queue priority.
type Priority int

const (
	PriorityNormal Priority = iota
	PriorityHigh
)

// Spec defines a mini-agent spawn request.
type Spec struct {
	ID           string
	SystemPrompt string
	Task         string
	Tools        []string
	ContextFiles []string
	Model        string
	Thinking     provider.Thinking
	Timeout      time.Duration
	Priority     Priority
}

// Result is the outcome of a mini-agent run.
type Result struct {
	Output  string
	Error   error
	Usage   provider.Usage
	Elapsed time.Duration
}

// Spawner launches and manages mini-agents.
type Spawner struct {
	provider provider.Provider
	tools    *tool.Registry
	pool     *Pool
	queue    *Queue
}

// NewSpawner creates a mini-agent spawner.
func NewSpawner(p provider.Provider, tools *tool.Registry, maxConcurrent int) *Spawner {
	return &Spawner{
		provider: p,
		tools:    tools,
		pool:     NewPool(maxConcurrent),
		queue:    NewQueue(),
	}
}

// Spawn launches a mini-agent and returns a handle.
func (s *Spawner) Spawn(ctx context.Context, spec Spec) (*Handle, error) {
	panic("not implemented")
}

// SpawnBatch launches multiple mini-agents concurrently.
func (s *Spawner) SpawnBatch(ctx context.Context, specs []Spec) ([]*Handle, error) {
	panic("not implemented")
}
