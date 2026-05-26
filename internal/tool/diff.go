package tool

import (
	"context"

	"github.com/mythosmystery/chef/internal/diff"
)

// DiffTool shows file changes.
type DiffTool struct {
	svc *diff.Service
}

// NewDiffTool creates a diff tool.
func NewDiffTool(svc *diff.Service) *DiffTool {
	return &DiffTool{svc: svc}
}

func (t *DiffTool) Name() string        { return "diff" }
func (t *DiffTool) Description() string { return "Show file changes" }
func (t *DiffTool) Parameters() map[string]any {
	return map[string]any{"path": map[string]any{"type": "string"}, "mode": map[string]any{"type": "string"}}
}
func (t *DiffTool) Call(ctx context.Context, args map[string]any) (Result, error) {
	panic("not implemented")
}
