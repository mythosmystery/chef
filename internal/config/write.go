package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// GlobalConfigPath returns the path to the global config file.
func GlobalConfigPath() (string, error) {
	dir, err := ChefDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

// ProjectConfigPath returns the path to the project config file.
func ProjectConfigPath(projectRoot string) string {
	return filepath.Join(projectRoot, ".chef", "config.json")
}

// WriteFile writes cfg to path as indented JSON, creating parent directories.
// The write is atomic (temp file + rename).
func WriteFile(path string, cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("config: marshal: %w", err)
	}
	data = append(data, '\n')

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("config: mkdir %s: %w", dir, err)
	}

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return fmt.Errorf("config: write %s: %w", tmp, err)
	}
	if err := os.Rename(tmp, path); err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf("config: rename %s: %w", path, err)
	}
	return nil
}
