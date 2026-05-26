package agent

// MergeResult collects validated output from sub-agents.
type MergeResult struct {
	AgentID string
	Output  string
	Diffs   []FileDiff
	Valid   bool
	Error   string
}

// FileDiff describes changes from a sub-agent.
type FileDiff struct {
	Path string
	Diff string
}

// Merger collects and validates sub-agent results before applying changes.
type Merger struct{}

// NewMerger creates a diff merger.
func NewMerger() *Merger {
	return &Merger{}
}

// Collect gathers results from completed sub-agents.
func (m *Merger) Collect(results []MergeResult) ([]FileDiff, error) {
	panic("not implemented")
}

// Validate checks sub-agent output against plan expectations.
func (m *Merger) Validate(result MergeResult) error {
	panic("not implemented")
}
