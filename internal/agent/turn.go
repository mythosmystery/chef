package agent

import "context"

// TurnResult is the outcome of one agent turn.
type TurnResult struct {
	Reply      string
	ToolCalls  int
	Compacted  bool
	StopReason string
}

// RunTurn executes prompt → tools → reply for one user message.
func (a *Agent) RunTurn(ctx context.Context, prompt string) (TurnResult, error) {
	panic("not implemented")
}
