package tui

import "github.com/mythosmystery/chef/internal/command"

// CommandsPopupModel renders / slash command completion suggestions.
type CommandsPopupModel struct {
	Visible  bool
	Commands []command.Command
	Selected int
	Width    int
}

// NewCommandsPopup creates a slash command completion popup.
func NewCommandsPopup() CommandsPopupModel {
	return CommandsPopupModel{}
}

// View renders the command completion popup.
func (p CommandsPopupModel) View() string {
	panic("not implemented")
}

// Show sets commands and makes the popup visible.
func (p CommandsPopupModel) Show(commands []command.Command) CommandsPopupModel {
	p.Visible = len(commands) > 0
	p.Commands = commands
	p.Selected = 0
	return p
}

// Hide clears and hides the popup.
func (p CommandsPopupModel) Hide() CommandsPopupModel {
	p.Visible = false
	p.Commands = nil
	p.Selected = 0
	return p
}
