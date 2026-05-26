package tool

import (
	"context"

	"github.com/mythosmystery/chef/internal/config"
)

// BashTool executes shell commands.
type BashTool struct {
	cfg config.BashConfig
}

// NewBashTool creates a bash tool.
func NewBashTool(cfg config.BashConfig) *BashTool {
	return &BashTool{cfg: cfg}
}

func (t *BashTool) Name() string        { return "bash" }
func (t *BashTool) Description() string { return "Execute shell commands" }
func (t *BashTool) Parameters() map[string]any {
	return map[string]any{"command": map[string]any{"type": "string"}}
}
func (t *BashTool) Call(ctx context.Context, args map[string]any) (Result, error) {
	panic("not implemented")
}
