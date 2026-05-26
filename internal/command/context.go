package command

import "context"

// RegisterContextCommands adds /context subcommands.
func RegisterContextCommands(r *Registry, handlers map[string]Handler) {
	for name, handler := range handlers {
		r.Register(Command{
			Name:        "context " + name,
			Description: "Project context operation: " + name,
			Usage:       "/context " + name,
			Handler:     handler,
		})
	}
}

// ContextQueryHandler handles /context query.
func ContextQueryHandler(ctx context.Context, args []string) error {
	panic("not implemented")
}

// ContextScanHandler handles /context scan.
func ContextScanHandler(ctx context.Context, args []string) error {
	panic("not implemented")
}

// ContextBudgetHandler handles /context budget.
func ContextBudgetHandler(ctx context.Context, args []string) error {
	panic("not implemented")
}

// ContextRefreshHandler handles /context refresh.
func ContextRefreshHandler(ctx context.Context, args []string) error {
	panic("not implemented")
}

// ContextRebuildHandler handles /context rebuild.
func ContextRebuildHandler(ctx context.Context, args []string) error {
	panic("not implemented")
}
