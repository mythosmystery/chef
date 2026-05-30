package workspace

import (
	"os"
	"path/filepath"
)

// FindGitRoot walks from start upward until a .git directory or file is found.
// Returns the absolute repo root, or the absolute start path if no git repo exists.
func FindGitRoot(start string) (string, error) {
	abs, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}

	dir := abs
	for {
		gitPath := filepath.Join(dir, ".git")
		if _, err := os.Stat(gitPath); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return abs, nil
		}
		dir = parent
	}
}
