package tool

import (
	"context"

	"github.com/mythosmystery/chef/internal/workspace"
)

// GrepTool searches file contents by regex.
type GrepTool struct {
	ws *workspace.Workspace
}

// NewGrepTool creates a grep tool.
func NewGrepTool(ws *workspace.Workspace) *GrepTool {
	return &GrepTool{ws: ws}
}

func (t *GrepTool) Name() string        { return "grep" }
func (t *GrepTool) Description() string { return "Search file contents by regex" }
func (t *GrepTool) Parameters() map[string]any {
	return map[string]any{"pattern": map[string]any{"type": "string"}, "path": map[string]any{"type": "string"}}
}
func (t *GrepTool) Call(ctx context.Context, args map[string]any) (Result, error) {
	panic("not implemented")
}
