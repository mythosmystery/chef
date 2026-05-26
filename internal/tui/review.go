package tui

// ReviewChoice is the user's decision on a review pager.
type ReviewChoice string

const (
	ReviewApprove ReviewChoice = "approve"
	ReviewRetry   ReviewChoice = "retry"
	ReviewSkip    ReviewChoice = "skip"
)

// ReviewModel is a generic Approve/Retry/Skip pager.
type ReviewModel struct {
	Title   string
	Content string
	Index   int
	Total   int
	Width   int
	Height  int
}

// NewReviewModel creates a review pager.
func NewReviewModel(title, content string, index, total int) ReviewModel {
	return ReviewModel{
		Title:   title,
		Content: content,
		Index:   index,
		Total:   total,
	}
}

// View renders the review pager.
func (r ReviewModel) View() string {
	panic("not implemented")
}
