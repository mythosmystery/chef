package config

import (
	"os"
	"path/filepath"
	"strings"
)

// ExpandPath expands a leading ~ to the user home directory.
func ExpandPath(path string) (string, error) {
	if path == "" {
		return "", nil
	}
	if path == "~" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return home, nil
	}
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[2:]), nil
	}
	return path, nil
}

// Resolve expands and absolutizes path fields against projectRoot.
func (c *Config) Resolve(projectRoot string) error {
	expanded, err := ExpandPath(c.Session.Dir)
	if err != nil {
		return err
	}
	if expanded != "" {
		if !filepath.IsAbs(expanded) {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			expanded = filepath.Join(home, expanded)
		}
		c.Session.Dir = filepath.Clean(expanded)
	}

	ctxDir := c.ContextFiles.Dir
	if ctxDir == "" {
		ctxDir = ".chef"
	}
	if filepath.IsAbs(ctxDir) {
		c.ContextFiles.Dir = filepath.Clean(ctxDir)
	} else {
		c.ContextFiles.Dir = filepath.Clean(filepath.Join(projectRoot, ctxDir))
	}
	return nil
}

// ChefDir returns the global chef config directory (~/.chef or CHEF_DIR).
func ChefDir() (string, error) {
	dir, err := chefDirRaw()
	if err != nil {
		return "", err
	}
	return ExpandPath(dir)
}

func chefDirRaw() (string, error) {
	if v := os.Getenv("CHEF_DIR"); v != "" {
		return v, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".chef"), nil
}
