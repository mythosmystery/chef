package tool

import (
	"context"
	"fmt"
)

// Registry holds registered tools and enforces allowlists.
type Registry struct {
	tools map[string]Tool
}

// NewRegistry creates an empty tool registry.
func NewRegistry() *Registry {
	return &Registry{tools: make(map[string]Tool)}
}

// Register adds a tool to the registry.
func (r *Registry) Register(t Tool) {
	r.tools[t.Name()] = t
}

// Get returns a tool by name.
func (r *Registry) Get(name string) (Tool, error) {
	t, ok := r.tools[name]
	if !ok {
		return nil, fmt.Errorf("unknown tool: %s", name)
	}
	return t, nil
}

// List returns tools filtered by allowlist (nil = all).
func (r *Registry) List(allowlist []string) []Tool {
	panic("not implemented")
}

// Definitions returns provider tool definitions for the allowlist.
func (r *Registry) Definitions(allowlist []string) []map[string]any {
	panic("not implemented")
}

// Invoke executes a tool call.
func (r *Registry) Invoke(ctx context.Context, call Call) (Result, error) {
	t, err := r.Get(call.Name)
	if err != nil {
		return Result{}, err
	}
	return t.Call(ctx, call.Args)
}
