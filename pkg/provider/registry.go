package provider

import "fmt"

// Registry maps provider names to constructors.
type Registry struct {
	providers map[string]func() (Provider, error)
}

// NewRegistry creates an empty provider registry.
func NewRegistry() *Registry {
	return &Registry{providers: make(map[string]func() (Provider, error))}
}

// Register adds a provider constructor by name.
func (r *Registry) Register(name string, ctor func() (Provider, error)) {
	r.providers[name] = ctor
}

// Get returns a provider by name.
func (r *Registry) Get(name string) (Provider, error) {
	ctor, ok := r.providers[name]
	if !ok {
		return nil, fmt.Errorf("unknown provider: %s", name)
	}
	return ctor()
}
