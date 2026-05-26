package tui

import "github.com/mythosmystery/chef/internal/refs"

// RefsPopupModel renders @ file reference completion suggestions.
type RefsPopupModel struct {
	Visible     bool
	Suggestions []refs.Completion
	Selected    int
	Width       int
}

// NewRefsPopup creates a reference completion popup.
func NewRefsPopup() RefsPopupModel {
	return RefsPopupModel{}
}

// View renders the completion popup.
func (p RefsPopupModel) View() string {
	panic("not implemented")
}

// Show sets suggestions and makes the popup visible.
func (p RefsPopupModel) Show(suggestions []refs.Completion) RefsPopupModel {
	p.Visible = len(suggestions) > 0
	p.Suggestions = suggestions
	p.Selected = 0
	return p
}

// Hide clears and hides the popup.
func (p RefsPopupModel) Hide() RefsPopupModel {
	p.Visible = false
	p.Suggestions = nil
	p.Selected = 0
	return p
}
