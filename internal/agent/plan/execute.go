package plan

import "context"

// Executor runs plan steps according to the dependency DAG.
type Executor struct{}

// NewExecutor creates a plan executor.
func NewExecutor() *Executor {
	return &Executor{}
}

// Execute runs all pending steps, delegating parallel work to sub-agents.
func (e *Executor) Execute(ctx context.Context, p *Plan) error {
	panic("not implemented")
}

// Resume continues execution from the first incomplete step.
func (e *Executor) Resume(ctx context.Context, p *Plan) error {
	panic("not implemented")
}
