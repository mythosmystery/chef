// Package tool implements the agent tool registry and built-in tools.
package tool

import "context"

// Tool is an agent-callable capability.
type Tool interface {
	Name() string
	Description() string
	Parameters() map[string]any
	Call(ctx context.Context, args map[string]any) (Result, error)
}

// Result is the output of a tool invocation.
type Result struct {
	Output string
	Error  error
}

// Call represents a tool invocation request.
type Call struct {
	Name string
	Args map[string]any
}
