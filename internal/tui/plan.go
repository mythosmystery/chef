package tui

import (
	"github.com/mythosmystery/chef/internal/agent/plan"
)

// PlanModel renders plan mode UI: review, execution, and failure prompts.
type PlanModel struct {
	Active  bool
	Plan    *plan.Plan
	Width   int
	Height  int
}

// NewPlanModel creates a plan UI model.
func NewPlanModel() PlanModel {
	return PlanModel{}
}

// View renders the plan panel.
func (p PlanModel) View() string {
	panic("not implemented")
}

// SetPlan updates the displayed plan.
func (p PlanModel) SetPlan(pl *plan.Plan) PlanModel {
	p.Plan = pl
	p.Active = pl != nil
	return p
}
