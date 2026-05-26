package agent

import (
	"context"

	"github.com/mythosmystery/chef/pkg/provider"
)

// StreamConsumer handles provider stream events during a turn.
type StreamConsumer struct {
	agent *Agent
}

// Consume reads stream events until done or error.
func (c *StreamConsumer) Consume(ctx context.Context, events <-chan provider.StreamEvent) error {
	panic("not implemented")
}
