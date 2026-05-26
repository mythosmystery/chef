package tui

import "github.com/charmbracelet/lipgloss"

var (
	colorBg        = lipgloss.Color("#1a1b26")
	colorSurface   = lipgloss.Color("#24283b")
	colorBorder    = lipgloss.Color("#414868")
	colorText      = lipgloss.Color("#c0caf5")
	colorFaint     = lipgloss.Color("#565f89")
	colorAccent    = lipgloss.Color("#7aa2f7")
	colorUser      = lipgloss.Color("#9ece6a")
	colorAssistant = lipgloss.Color("#bb9af7")
)

var (
	Header = lipgloss.NewStyle().
		Foreground(colorAccent).
		Bold(true).
		Padding(0, 1)

	Footer = lipgloss.NewStyle().
		Foreground(colorFaint).
		Padding(0, 1)

	Faint = lipgloss.NewStyle().
		Foreground(colorFaint)

	UserPrefix = lipgloss.NewStyle().
			Foreground(colorUser).
			Bold(true)

	UserBody = lipgloss.NewStyle().
			Foreground(colorText).
			PaddingLeft(2)

	AssistantPrefix = lipgloss.NewStyle().
			Foreground(colorAssistant).
			Bold(true)

	AssistantBody = lipgloss.NewStyle().
			Foreground(colorText).
			PaddingLeft(2)

	InputFrame = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(0, 1)

	LogFrame = lipgloss.NewStyle().
			Padding(0, 1)
)
