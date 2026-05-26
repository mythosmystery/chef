// Package workspace provides filesystem access and path safety for chef tools.
package workspace

import (
	"os"
	"path/filepath"
)

// Workspace is the project root and safety boundaries.
type Workspace struct {
	Root       string
	ContextDir string
	SessionDir string
}

// New creates a workspace for the given project root.
func New(root, contextDir, sessionDir string) (*Workspace, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	return &Workspace{
		Root:       abs,
		ContextDir: contextDir,
		SessionDir: sessionDir,
	}, nil
}

// CWD returns the current working directory or root if unavailable.
func CWD() (string, error) {
	return os.Getwd()
}
