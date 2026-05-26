package provider

// Role identifies message authors in provider messages.
type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
	RoleTool      Role = "tool"
)

// Message is a provider-agnostic chat message.
type Message struct {
	Role    Role
	Content []Part
}

// Part is a segment of message content.
type Part struct {
	Type     string
	Text     string
	ToolCall *ToolCall
	ToolID   string
	Name     string
}
