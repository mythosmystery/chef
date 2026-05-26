package tui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	headerLines = 1
	footerLines = 1
)

type Model struct {
	input  InputModel
	log    LogModel
	width  int
	height int
	cwd    string
}

func New() Model {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}

	return Model{
		input: NewInput(),
		log:   NewLog(),
		cwd:   cwd,
	}
}

func (m Model) Init() tea.Cmd {
	return m.input.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.input.SetWidth(m.contentWidth())
		m.log = m.log.SetSize(m.contentWidth(), m.logHeight())
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		}

		var cmd tea.Cmd
		var handled bool
		m.input, cmd, handled = m.input.HandleKey(msg)
		if handled {
			return m, cmd
		}

	case submitMsg:
		m.log = m.log.Append(NewUserMessage(msg.content))
		return m, nil
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) contentWidth() int {
	if m.width <= 0 {
		return 80
	}
	return m.width
}

func (m Model) logHeight() int {
	h := m.height - headerLines - inputHeight - footerLines
	if h < 1 {
		return 1
	}
	return h
}
