package provider

// StreamEventKind identifies stream event types.
type StreamEventKind string

const (
	EventTextDelta     StreamEventKind = "text_delta"
	EventThinkingDelta StreamEventKind = "thinking_delta"
	EventToolCallStart StreamEventKind = "tool_call_start"
	EventToolCallDelta StreamEventKind = "tool_call_delta"
	EventToolCallEnd   StreamEventKind = "tool_call_end"
	EventUsage         StreamEventKind = "usage"
	EventDone          StreamEventKind = "done"
	EventError         StreamEventKind = "error"
)

// StreamEvent is one event from a streaming completion.
type StreamEvent struct {
	Kind     StreamEventKind
	Text     string
	ToolCall ToolCall
	Usage    Usage
	Error    error
}

// Usage reports token usage for a completion.
type Usage struct {
	InputTokens  int
	OutputTokens int
}
