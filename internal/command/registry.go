package command

import "fmt"

// Registry holds available slash commands.
type Registry struct {
	commands map[string]Command
}

// NewRegistry creates an empty command registry.
func NewRegistry() *Registry {
	return &Registry{commands: make(map[string]Command)}
}

// Register adds a command by name.
func (r *Registry) Register(cmd Command) {
	r.commands[cmd.Name] = cmd
}

// Get returns a command by name.
func (r *Registry) Get(name string) (Command, error) {
	cmd, ok := r.commands[name]
	if !ok {
		return Command{}, fmt.Errorf("unknown command: /%s", name)
	}
	return cmd, nil
}

// All returns all registered commands.
func (r *Registry) All() []Command {
	out := make([]Command, 0, len(r.commands))
	for _, cmd := range r.commands {
		out = append(out, cmd)
	}
	return out
}
