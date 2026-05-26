package spawn

import (
	"context"
	"sync"
)

// HandleStatus tracks mini-agent lifecycle state.
type HandleStatus string

const (
	HandleQueued   HandleStatus = "queued"
	HandleRunning  HandleStatus = "running"
	HandleDone     HandleStatus = "done"
	HandleFailed   HandleStatus = "failed"
	HandleTimedOut HandleStatus = "timed_out"
	HandleCanceled HandleStatus = "canceled"
)

// Handle tracks a running mini-agent.
type Handle struct {
	ID     string
	Spec   Spec
	Status HandleStatus
	Result Result

	mu     sync.Mutex
	cancel context.CancelFunc
	done   chan struct{}
}

// Cancel requests cancellation of the mini-agent.
func (h *Handle) Cancel() {
	panic("not implemented")
}

// Wait blocks until the mini-agent completes.
func (h *Handle) Wait(ctx context.Context) (Result, error) {
	panic("not implemented")
}

// Snapshot returns a copy of current status for TUI display.
func (h *Handle) Snapshot() Handle {
	panic("not implemented")
}
