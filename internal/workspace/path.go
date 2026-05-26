package workspace

import "path/filepath"

// Normalize resolves and cleans a path relative to the workspace root.
func (w *Workspace) Normalize(path string) (string, error) {
	if filepath.IsAbs(path) {
		return filepath.Clean(path), nil
	}
	return filepath.Join(w.Root, filepath.Clean(path)), nil
}

// Rel returns path relative to workspace root.
func (w *Workspace) Rel(absPath string) (string, error) {
	return filepath.Rel(w.Root, absPath)
}
