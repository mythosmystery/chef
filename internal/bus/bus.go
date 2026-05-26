// Package bus provides the event bus between the agent and TUI.
package bus

import "sync"

// Bus routes events between async agent work and the synchronous TUI.
type Bus struct {
	mu       sync.RWMutex
	listener func(Event)
}

// New creates an event bus.
func New() *Bus {
	return &Bus{}
}

// Subscribe registers a listener for agent events.
func (b *Bus) Subscribe(fn func(Event)) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.listener = fn
}

// Publish sends an event to the subscribed listener.
func (b *Bus) Publish(e Event) {
	b.mu.RLock()
	fn := b.listener
	b.mu.RUnlock()
	if fn != nil {
		fn(e)
	}
}
