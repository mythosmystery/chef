package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type LogModel struct {
	viewport viewport.Model
	messages []Message
	width    int
}

func NewLog() LogModel {
	return LogModel{viewport: viewport.New(0, 0)}
}

func (m LogModel) Init() tea.Cmd {
	return nil
}

func (m LogModel) Update(msg tea.Msg) (LogModel, tea.Cmd) {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m LogModel) Append(msg Message) LogModel {
	m.messages = append(m.messages, msg)
	m.refreshContent()
	return m
}

func (m LogModel) SetSize(width, height int) LogModel {
	m.width = width
	m.viewport.Width = width
	m.viewport.Height = height
	m.refreshContent()
	return m
}

func (m *LogModel) refreshContent() {
	var b strings.Builder
	for _, msg := range m.messages {
		b.WriteString(renderMessage(msg, m.width))
		b.WriteByte('\n')
	}
	m.viewport.SetContent(b.String())
	m.viewport.GotoBottom()
}

func renderMessage(msg Message, width int) string {
	ts := msg.At.Format("15:04")

	switch msg.Role {
	case RoleUser:
		header := UserPrefix.Render(fmt.Sprintf("you  %s", ts))
		body := UserBody.Render("> " + msg.Content)
		return header + "\n" + body
	case RoleAssistant:
		header := AssistantPrefix.Render(fmt.Sprintf("chef  %s", ts))
		body := AssistantBody.Render(msg.Content)
		return header + "\n" + body
	case RoleSystem:
		return Faint.Render(msg.Content)
	default:
		return msg.Content
	}
}

func (m LogModel) View() string {
	if len(m.messages) == 0 {
		empty := Faint.Render("No messages yet. Type below and press Enter to send.")
		return LogFrame.Render(empty)
	}
	return LogFrame.Render(m.viewport.View())
}

// NewUserMessage creates a user message with the current timestamp.
func NewUserMessage(content string) Message {
	return Message{
		Role:    RoleUser,
		Content: content,
		At:      time.Now(),
	}
}
