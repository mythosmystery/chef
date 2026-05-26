package command

import "context"

// RegisterSessionCommands adds session-related slash commands.
func RegisterSessionCommands(r *Registry) {
	r.Register(Command{Name: "session", Description: "Show session info", Usage: "/session", Handler: SessionHandler})
	r.Register(Command{Name: "resume", Description: "Browse past sessions", Usage: "/resume", Handler: ResumeHandler})
	r.Register(Command{Name: "new", Description: "Start new session", Usage: "/new", Handler: NewSessionHandler})
}

func SessionHandler(ctx context.Context, args []string) error {
	panic("not implemented")
}

func ResumeHandler(ctx context.Context, args []string) error {
	panic("not implemented")
}

func NewSessionHandler(ctx context.Context, args []string) error {
	panic("not implemented")
}
