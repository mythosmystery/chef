package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

const inputHeight = 4

type submitMsg struct {
	content string
}

type InputModel struct {
	area textarea.Model
}

func NewInput() InputModel {
	ta := textarea.New()
	ta.Placeholder = "type a message..."
	ta.ShowLineNumbers = false
	ta.CharLimit = 8000
	ta.SetHeight(inputHeight - 2) // account for border + padding
	ta.Focus()

	// Plain Enter is handled by the parent; Shift+Enter inserts a newline.
	ta.KeyMap.InsertNewline = key.NewBinding(
		key.WithKeys("shift+enter"),
		key.WithHelp("shift+enter", "insert newline"),
	)

	return InputModel{area: ta}
}

func (m InputModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m InputModel) Update(msg tea.Msg) (InputModel, tea.Cmd) {
	var cmd tea.Cmd
	m.area, cmd = m.area.Update(msg)
	return m, cmd
}

func (m InputModel) Value() string {
	return m.area.Value()
}

func (m InputModel) Clear() {
	m.area.Reset()
}

func (m InputModel) SetWidth(width int) {
	m.area.SetWidth(width - 4) // border + horizontal padding
}

func (m InputModel) View() string {
	return InputFrame.Render(m.area.View())
}

func (m InputModel) Submit() (InputModel, tea.Cmd) {
	content := strings.TrimSpace(m.area.Value())
	if content == "" {
		return m, nil
	}
	m.area.Reset()
	return m, func() tea.Msg { return submitMsg{content: content} }
}

func (m InputModel) HandleKey(msg tea.KeyMsg) (InputModel, tea.Cmd, bool) {
	switch msg.String() {
	case "enter":
		m, cmd := m.Submit()
		return m, cmd, true
	case "shift+enter":
		var cmd tea.Cmd
		m.area, cmd = m.area.Update(msg)
		return m, cmd, true
	}
	return m, nil, false
}
