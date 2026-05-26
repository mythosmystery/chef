package spawn

import "context"

// Pool limits concurrent mini-agent goroutines.
type Pool struct {
	sem chan struct{}
}

// NewPool creates a concurrency pool with max slots.
func NewPool(max int) *Pool {
	if max <= 0 {
		max = 1
	}
	return &Pool{sem: make(chan struct{}, max)}
}

// Acquire blocks until a slot is available or ctx is cancelled.
func (p *Pool) Acquire(ctx context.Context) error {
	select {
	case p.sem <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Release returns a slot to the pool.
func (p *Pool) Release() {
	select {
	case <-p.sem:
	default:
	}
}
