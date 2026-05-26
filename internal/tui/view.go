package tui

import (
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	return strings.Join([]string{
		renderHeader(m.cwd, m.contentWidth()),
		m.log.View(),
		m.input.View(),
		renderFooter(m.contentWidth()),
	}, "\n")
}

func renderHeader(cwd string, width int) string {
	left := Header.Render("chef")
	sep := Faint.Render(" · ")
	path := Faint.Render(shortenPath(cwd))

	content := left + sep + path
	if lipgloss.Width(content) > width {
		content = lipgloss.NewStyle().Width(width).Render(content)
	}
	return content
}

func renderFooter(width int) string {
	hints := Faint.Render("enter send  ·  shift+enter newline  ·  ctrl+c quit")
	if lipgloss.Width(hints) > width {
		hints = lipgloss.NewStyle().Width(width).Render(hints)
	}
	return Footer.Render(hints)
}

func shortenPath(path string) string {
	home, err := os.UserHomeDir()
	if err == nil && strings.HasPrefix(path, home) {
		return "~" + strings.TrimPrefix(path, home)
	}
	return path
}
