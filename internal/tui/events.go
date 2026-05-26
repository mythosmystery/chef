package tui

import (
	"time"

	"github.com/mythosmystery/chef/internal/bus"
)

// AgentEventMsg wraps a bus event for the TUI update loop.
type AgentEventMsg struct {
	Event bus.Event
}

// AssistantDeltaMsg carries streamed assistant text.
type AssistantDeltaMsg struct {
	Text string
	At   time.Time
}

// ToolCallMsg carries a tool invocation for display.
type ToolCallMsg struct {
	Name string
	Args map[string]any
	At   time.Time
}

// ToolResultMsg carries a tool result for display.
type ToolResultMsg struct {
	Name   string
	Output string
	Error  string
	At     time.Time
}

// AgentProgressMsg carries mini-agent status for the progress widget.
type AgentProgressMsg struct {
	AgentID string
	Task    string
	Status  string
	Detail  string
}

// PlanUpdateMsg carries plan execution state changes.
type PlanUpdateMsg struct {
	PlanID string
	StepID string
	Status string
}

// ContextUpdateMsg carries context maintenance summaries.
type ContextUpdateMsg struct {
	Summary string
}

// AgentErrorMsg carries a surfaced agent error.
type AgentErrorMsg struct {
	Message string
	At      time.Time
}

// AgentDoneMsg signals completion of an agent turn.
type AgentDoneMsg struct {
	At time.Time
}

// FromBusEvent converts a bus event to a tea.Msg if applicable.
func FromBusEvent(e bus.Event) teaMsg {
	switch e.Kind {
	case bus.EventAssistantDelta:
		if p, ok := e.Payload.(bus.AssistantDelta); ok {
			return AssistantDeltaMsg{Text: p.Text, At: e.At}
		}
	case bus.EventToolCall:
		if p, ok := e.Payload.(bus.ToolCallEvent); ok {
			return ToolCallMsg{Name: p.Name, Args: p.Args, At: e.At}
		}
	case bus.EventToolResult:
		if p, ok := e.Payload.(bus.ToolResultEvent); ok {
			return ToolResultMsg{Name: p.Name, Output: p.Output, Error: p.Error, At: e.At}
		}
	case bus.EventAgentProgress:
		if p, ok := e.Payload.(bus.AgentProgressEvent); ok {
			return AgentProgressMsg{AgentID: p.AgentID, Task: p.Task, Status: p.Status, Detail: p.Detail}
		}
	case bus.EventPlanUpdate:
		if p, ok := e.Payload.(bus.PlanUpdateEvent); ok {
			return PlanUpdateMsg{PlanID: p.PlanID, StepID: p.StepID, Status: p.Status}
		}
	case bus.EventContextUpdate:
		if p, ok := e.Payload.(bus.ContextUpdateEvent); ok {
			return ContextUpdateMsg{Summary: p.Summary}
		}
	case bus.EventError:
		if p, ok := e.Payload.(bus.ErrorEvent); ok {
			return AgentErrorMsg{Message: p.Message, At: e.At}
		}
	case bus.EventDone:
		return AgentDoneMsg{At: e.At}
	}
	return AgentEventMsg{Event: e}
}

// teaMsg is the union of agent-originated messages for the TUI.
type teaMsg any
