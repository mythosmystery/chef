package tui

// HeaderModel renders session name, model, and plan mode indicator.
type HeaderModel struct {
	SessionName string
	Model       string
	PlanTitle   string
	Width       int
}

// NewHeader creates a header model.
func NewHeader() HeaderModel {
	return HeaderModel{}
}

// View renders the header line.
func (h HeaderModel) View() string {
	return renderHeaderModel(h)
}

func renderHeaderModel(h HeaderModel) string {
	panic("not implemented")
}
