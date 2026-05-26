package command

import "context"

// RegisterMetaCommands adds model and quit commands.
func RegisterMetaCommands(r *Registry) {
	r.Register(Command{Name: "model", Description: "Switch model", Usage: "/model", Handler: ModelHandler})
	r.Register(Command{Name: "quit", Description: "Exit chef", Usage: "/quit", Handler: QuitHandler})
}

func ModelHandler(ctx context.Context, args []string) error {
	panic("not implemented")
}

func QuitHandler(ctx context.Context, args []string) error {
	panic("not implemented")
}

// DefaultRegistry returns a registry with all built-in commands registered.
func DefaultRegistry() *Registry {
	r := NewRegistry()
	RegisterPlanCommands(r, PlanHandler)
	RegisterContextCommands(r, map[string]Handler{
		"query":   ContextQueryHandler,
		"scan":    ContextScanHandler,
		"budget":  ContextBudgetHandler,
		"refresh": ContextRefreshHandler,
		"rebuild": ContextRebuildHandler,
	})
	RegisterSessionCommands(r)
	RegisterMetaCommands(r)
	return r
}
