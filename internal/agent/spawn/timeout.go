package spawn

import (
	"context"
	"time"
)

// WithTimeout wraps ctx with the spec timeout and returns a cancel func.
func WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	if timeout <= 0 {
		return parent, func() {}
	}
	return context.WithTimeout(parent, timeout)
}

// TimedOut reports whether err is a timeout.
func TimedOut(err error) bool {
	panic("not implemented")
}
