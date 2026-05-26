package tui

// InitPhase identifies context initialization UI phases.
type InitPhase string

const (
	InitProgress InitPhase = "progress"
	InitReview   InitPhase = "review"
)

// InitModel renders context init progress and per-file review.
type InitModel struct {
	Phase     InitPhase
	Tasks     []InitTaskStatus
	Files     []InitFileReview
	FileIndex int
	Width     int
	Height    int
}

// InitTaskStatus tracks one init mini-agent.
type InitTaskStatus struct {
	Path   string
	Status string
	Detail string
}

// InitFileReview is a generated context file awaiting user review.
type InitFileReview struct {
	Name    string
	Content string
}

// NewInitModel creates an init flow UI model.
func NewInitModel() InitModel {
	return InitModel{Phase: InitProgress}
}

// View renders the init flow screen.
func (m InitModel) View() string {
	panic("not implemented")
}
