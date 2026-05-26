package bus

import "time"

// EventKind identifies agent-to-TUI event types.
type EventKind string

const (
	EventAssistantDelta EventKind = "assistant_delta"
	EventToolCall       EventKind = "tool_call"
	EventToolResult     EventKind = "tool_result"
	EventAgentProgress  EventKind = "agent_progress"
	EventPlanUpdate     EventKind = "plan_update"
	EventContextUpdate  EventKind = "context_update"
	EventError          EventKind = "error"
	EventDone           EventKind = "done"
)

// Event is a message flowing from the agent to the TUI.
type Event struct {
	Kind      EventKind
	SessionID string
	At        time.Time
	Payload   any
}

// AssistantDelta carries streamed assistant text.
type AssistantDelta struct {
	Text string
}

// ToolCallEvent carries a tool invocation.
type ToolCallEvent struct {
	Name string
	Args map[string]any
}

// ToolResultEvent carries a tool result.
type ToolResultEvent struct {
	Name   string
	Output string
	Error  string
}

// AgentProgressEvent carries mini-agent status updates.
type AgentProgressEvent struct {
	AgentID string
	Task    string
	Status  string
	Detail  string
}

// PlanUpdateEvent carries plan execution state changes.
type PlanUpdateEvent struct {
	PlanID string
	StepID string
	Status string
}

// ContextUpdateEvent carries context maintenance summaries.
type ContextUpdateEvent struct {
	Summary string
}

// ErrorEvent carries a surfaced error.
type ErrorEvent struct {
	Message string
}
