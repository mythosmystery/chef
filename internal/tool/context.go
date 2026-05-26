package tool

import (
	"context"

	"github.com/mythosmystery/chef/internal/projctx"
)

// ContextTool operates on .chef context files.
type ContextTool struct {
	mgr *projctx.Manager
}

// NewContextTool creates a context tool.
func NewContextTool(mgr *projctx.Manager) *ContextTool {
	return &ContextTool{mgr: mgr}
}

func (t *ContextTool) Name() string        { return "context" }
func (t *ContextTool) Description() string { return "Query and maintain project context files" }
func (t *ContextTool) Parameters() map[string]any {
	return map[string]any{"operation": map[string]any{"type": "string"}, "args": map[string]any{"type": "object"}}
}
func (t *ContextTool) Call(ctx context.Context, args map[string]any) (Result, error) {
	panic("not implemented")
}
