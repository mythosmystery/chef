// Package refs parses and resolves @ file references in user input.
package refs

import "github.com/mythosmystery/chef/internal/workspace"

// Kind identifies reference types.
type Kind string

const (
	KindFile      Kind = "file"
	KindLineRange Kind = "line_range"
	KindDirectory Kind = "directory"
	KindContext   Kind = "context"
)

// Reference is a parsed @ attachment.
type Reference struct {
	Kind    Kind
	Path    string
	Start   int
	End     int
	Raw     string
}

// Resolver resolves @ references to attachable content.
type Resolver struct {
	ws *workspace.Workspace
}

// NewResolver creates a reference resolver.
func NewResolver(ws *workspace.Workspace) *Resolver {
	return &Resolver{ws: ws}
}
