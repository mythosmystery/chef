package tool

import (
	"context"

	"github.com/mythosmystery/chef/internal/workspace"
)

// ReadTool reads file contents.
type ReadTool struct {
	ws *workspace.Workspace
}

// NewReadTool creates a read tool.
func NewReadTool(ws *workspace.Workspace) *ReadTool {
	return &ReadTool{ws: ws}
}

func (t *ReadTool) Name() string        { return "read" }
func (t *ReadTool) Description() string { return "Read file contents" }
func (t *ReadTool) Parameters() map[string]any {
	return map[string]any{"path": map[string]any{"type": "string"}}
}
func (t *ReadTool) Call(ctx context.Context, args map[string]any) (Result, error) {
	panic("not implemented")
}
