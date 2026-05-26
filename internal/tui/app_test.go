package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func applyCmd(m Model, cmd tea.Cmd) Model {
	if cmd == nil {
		return m
	}
	msg := cmd()
	if msg == nil {
		return m
	}
	updated, _ := m.Update(msg)
	return updated.(Model)
}

func TestEnterSubmitsAndClearsInput(t *testing.T) {
	m := New()
	updated, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	m = updated.(Model)

	for _, r := range "hello" {
		var cmd tea.Cmd
		updated, cmd = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}})
		m = applyCmd(updated.(Model), cmd)
	}

	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected submit cmd on enter")
	}
	m = applyCmd(updated.(Model), cmd)

	view := m.View()
	if !strings.Contains(view, "> hello") {
		t.Fatalf("expected submitted message in view, got:\n%s", view)
	}
	if m.input.Value() != "" {
		t.Fatalf("expected cleared input, value=%q", m.input.Value())
	}
}

func TestLogAppendRendersMessage(t *testing.T) {
	log := NewLog()
	log = log.SetSize(80, 10)
	log = log.Append(NewUserMessage("hello"))

	view := log.View()
	if !strings.Contains(view, "hello") {
		t.Fatalf("expected hello in log view, got:\n%s", view)
	}
}
