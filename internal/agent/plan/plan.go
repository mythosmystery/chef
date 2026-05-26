// Package plan implements plan generation and DAG-based execution.
package plan

// Status is a plan step execution status.
type Status string

const (
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusDone    Status = "done"
	StatusFailed  Status = "failed"
	StatusSkipped Status = "skipped"
)

// Step is one unit of work in a plan.
type Step struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Files       []string `json:"files"`
	DependsOn   []string `json:"dependsOn"`
	Status      Status   `json:"status"`
	AssignedTo  string   `json:"assignedTo"`
	Output      string   `json:"output"`
}

// Plan is a structured, resumable execution plan.
type Plan struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Steps       []Step `json:"steps"`
}
