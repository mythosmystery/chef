package tool

import (
	"context"

	"github.com/mythosmystery/chef/internal/diff"
	"github.com/mythosmystery/chef/internal/workspace"
)

// EditTool performs precise text replacement edits.
type EditTool struct {
	ws      *workspace.Workspace
	tracker *diff.Tracker
}

// NewEditTool creates an edit tool.
func NewEditTool(ws *workspace.Workspace, tracker *diff.Tracker) *EditTool {
	return &EditTool{ws: ws, tracker: tracker}
}

func (t *EditTool) Name() string        { return "edit" }
func (t *EditTool) Description() string { return "Replace exact text in a file" }
func (t *EditTool) Parameters() map[string]any {
	return map[string]any{
		"path":    map[string]any{"type": "string"},
		"oldText": map[string]any{"type": "string"},
		"newText": map[string]any{"type": "string"},
	}
}
func (t *EditTool) Call(ctx context.Context, args map[string]any) (Result, error) {
	panic("not implemented")
}
