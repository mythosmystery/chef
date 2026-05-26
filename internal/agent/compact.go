package agent

import "github.com/mythosmystery/chef/internal/session"

// Compactor summarizes old session messages when thresholds are reached.
type Compactor struct {
	agent *Agent
}

// NewCompactor creates a session compactor.
func NewCompactor(a *Agent) *Compactor {
	return &Compactor{agent: a}
}

// MaybeCompact compacts session history if token or turn thresholds are met.
func (c *Compactor) MaybeCompact(messages []session.Message) ([]session.Message, bool, error) {
	panic("not implemented")
}

// FinalSummary generates the end-of-session summary for resume.
func (c *Compactor) FinalSummary(messages []session.Message) (string, error) {
	panic("not implemented")
}
