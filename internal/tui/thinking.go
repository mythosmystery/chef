package tui

// ThinkingModel renders collapsible thinking blocks in the message log.
type ThinkingModel struct {
	Content   string
	Expanded  bool
	Width     int
}

// NewThinkingModel creates a thinking block.
func NewThinkingModel(content string) ThinkingModel {
	return ThinkingModel{Content: content}
}

// View renders the thinking block.
func (t ThinkingModel) View() string {
	panic("not implemented")
}

// Toggle expands or collapses the thinking block.
func (t ThinkingModel) Toggle() ThinkingModel {
	t.Expanded = !t.Expanded
	return t
}
