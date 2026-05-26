package plan

// FailureAction is the user's choice when a step fails.
type FailureAction string

const (
	FailureRetry   FailureAction = "retry"
	FailureSkip    FailureAction = "skip"
	FailureEdit    FailureAction = "edit"
	FailureCancel  FailureAction = "cancel"
)

// FailureHandler manages step failure retry and user prompts.
type FailureHandler struct {
	maxRetries int
}

// NewFailureHandler creates a failure handler.
func NewFailureHandler(maxRetries int) *FailureHandler {
	return &FailureHandler{maxRetries: maxRetries}
}

// Handle processes a failed step and returns the next action.
func (h *FailureHandler) Handle(p *Plan, step Step, err error, attempt int) (FailureAction, error) {
	panic("not implemented")
}
