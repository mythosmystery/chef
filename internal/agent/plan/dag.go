package plan

// DAG represents step dependencies and parallelizable groups.
type DAG struct {
	steps map[string]Step
}

// NewDAG builds a DAG from plan steps.
func NewDAG(steps []Step) (*DAG, error) {
	panic("not implemented")
}

// Ready returns steps whose dependencies are satisfied.
func (d *DAG) Ready(completed map[string]bool) []Step {
	panic("not implemented")
}

// ParallelGroups returns groups of steps that can run concurrently.
func (d *DAG) ParallelGroups() ([][]Step, error) {
	panic("not implemented")
}

// Validate checks for cycles and missing dependency references.
func (d *DAG) Validate() error {
	panic("not implemented")
}
