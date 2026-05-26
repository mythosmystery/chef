package tui

import "time"

const (
	RoleUser      = "user"
	RoleAssistant = "assistant"
	RoleSystem    = "system"
)

type Message struct {
	Role    string
	Content string
	At      time.Time
}
