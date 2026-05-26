package plan

import "context"

// ReviewAction is the user's decision during plan review.
type ReviewAction string

const (
	ReviewApprove ReviewAction = "approve"
	ReviewRevise  ReviewAction = "revise"
	ReviewCancel  ReviewAction = "cancel"
)

// Reviewer handles conversational plan review.
type Reviewer struct{}

// NewReviewer creates a plan reviewer.
func NewReviewer() *Reviewer {
	return &Reviewer{}
}

// Review presents a plan and applies user revision requests until approved.
func (r *Reviewer) Review(ctx context.Context, p *Plan, userMessage string) (*Plan, ReviewAction, error) {
	panic("not implemented")
}
