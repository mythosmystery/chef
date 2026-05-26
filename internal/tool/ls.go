package tool

import (
	"context"

	"github.com/mythosmystery/chef/internal/workspace"
)

// LSTool lists directory contents.
type LSTool struct {
	ws *workspace.Workspace
}

// NewLSTool creates an ls tool.
func NewLSTool(ws *workspace.Workspace) *LSTool {
	return &LSTool{ws: ws}
}

func (t *LSTool) Name() string        { return "ls" }
func (t *LSTool) Description() string { return "List directory contents" }
func (t *LSTool) Parameters() map[string]any {
	return map[string]any{"path": map[string]any{"type": "string"}}
}
func (t *LSTool) Call(ctx context.Context, args map[string]any) (Result, error) {
	panic("not implemented")
}
