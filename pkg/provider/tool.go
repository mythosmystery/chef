package provider

// ToolDef describes a tool available to the model.
type ToolDef struct {
	Name        string
	Description string
	Parameters  map[string]any
}

// ToolCall is a model-initiated tool invocation.
type ToolCall struct {
	ID   string
	Name string
	Args map[string]any
}

// ToolResult is the outcome of a tool call returned to the model.
type ToolResult struct {
	ToolCallID string
	Name       string
	Content    string
	IsError    bool
}
