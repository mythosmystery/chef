package tool

import (
	"context"

	"github.com/mythosmystery/chef/internal/workspace"
)

// WriteTool creates or overwrites files.
type WriteTool struct {
	ws *workspace.Workspace
}

// NewWriteTool creates a write tool.
func NewWriteTool(ws *workspace.Workspace) *WriteTool {
	return &WriteTool{ws: ws}
}

func (t *WriteTool) Name() string        { return "write" }
func (t *WriteTool) Description() string { return "Create or overwrite files" }
func (t *WriteTool) Parameters() map[string]any {
	return map[string]any{"path": map[string]any{"type": "string"}, "content": map[string]any{"type": "string"}}
}
func (t *WriteTool) Call(ctx context.Context, args map[string]any) (Result, error) {
	panic("not implemented")
}
