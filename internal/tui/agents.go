package tui

// AgentStatus tracks one mini-agent for the progress widget.
type AgentStatus struct {
	ID     string
	Task   string
	Status string
	Detail string
}

// AgentsModel renders mini-agent progress above the input editor.
type AgentsModel struct {
	Agents []AgentStatus
	Width  int
}

// NewAgentsModel creates a mini-agent progress widget.
func NewAgentsModel() AgentsModel {
	return AgentsModel{}
}

// View renders active mini-agent progress.
func (a AgentsModel) View() string {
	panic("not implemented")
}

// UpdateAgent upserts an agent status by ID.
func (a AgentsModel) UpdateAgent(status AgentStatus) AgentsModel {
	for i, existing := range a.Agents {
		if existing.ID == status.ID {
			a.Agents[i] = status
			return a
		}
	}
	a.Agents = append(a.Agents, status)
	return a
}

// RemoveAgent removes an agent by ID.
func (a AgentsModel) RemoveAgent(id string) AgentsModel {
	out := a.Agents[:0]
	for _, agent := range a.Agents {
		if agent.ID != id {
			out = append(out, agent)
		}
	}
	a.Agents = out
	return a
}
