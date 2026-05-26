package tool

import (
	"context"

	"github.com/mythosmystery/chef/internal/workspace"
)

// FindTool finds files by glob pattern.
type FindTool struct {
	ws *workspace.Workspace
}

// NewFindTool creates a find tool.
func NewFindTool(ws *workspace.Workspace) *FindTool {
	return &FindTool{ws: ws}
}

func (t *FindTool) Name() string        { return "find" }
func (t *FindTool) Description() string { return "Find files by glob pattern" }
func (t *FindTool) Parameters() map[string]any {
	return map[string]any{"pattern": map[string]any{"type": "string"}, "path": map[string]any{"type": "string"}}
}
func (t *FindTool) Call(ctx context.Context, args map[string]any) (Result, error) {
	panic("not implemented")
}
