// Package command parses and dispatches slash commands from user input.
package command

import "context"

// Handler executes a parsed slash command.
type Handler func(ctx context.Context, args []string) error

// Command describes a registered slash command.
type Command struct {
	Name        string
	Description string
	Usage       string
	Handler     Handler
}
