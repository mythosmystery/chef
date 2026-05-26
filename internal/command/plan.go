package command

import "context"

// RegisterPlanCommands adds /plan command handlers.
func RegisterPlanCommands(r *Registry, handler Handler) {
	r.Register(Command{
		Name:        "plan",
		Description: "Enter plan mode",
		Usage:       "/plan <prompt>",
		Handler:     handler,
	})
}

// PlanHandler is a stub plan command handler.
func PlanHandler(ctx context.Context, args []string) error {
	panic("not implemented")
}
