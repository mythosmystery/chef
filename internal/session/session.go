// Package session persists conversation history as JSONL with tree structure.
package session

import "time"

// Role identifies message authors.
type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
	RoleTool      Role = "tool"
)

// Part is a segment of message content.
type Part struct {
	Type    string `json:"type"`
	Content string `json:"content,omitempty"`
	Name    string `json:"name,omitempty"`
}

// Message is one conversation turn.
type Message struct {
	ID        string    `json:"id"`
	ParentID  string    `json:"parentId,omitempty"`
	Role      Role      `json:"role"`
	Parts     []Part    `json:"parts"`
	CreatedAt time.Time `json:"createdAt"`
}

// Session is a persisted conversation tree node.
type Session struct {
	ID         string    `json:"id"`
	ParentID   string    `json:"parentId,omitempty"`
	Title      string    `json:"title"`
	WorkingDir string    `json:"workingDir"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Summary    string    `json:"summary,omitempty"`
	PlanJSON   string    `json:"planJson,omitempty"`
}
