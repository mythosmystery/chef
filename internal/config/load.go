package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Load reads global and project config, merges them, and applies defaults.
func Load() (*Config, error) {
	cfg := Defaults()

	globalPath, err := globalConfigPath()
	if err != nil {
		return nil, err
	}
	if err := mergeFile(&cfg, globalPath); err != nil {
		return nil, err
	}

	projectPath := filepath.Join(".chef", "config.json")
	if err := mergeFile(&cfg, projectPath); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func mergeFile(cfg *Config, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, cfg)
}

func globalConfigPath() (string, error) {
	dir, err := chefDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

func chefDir() (string, error) {
	if v := os.Getenv("CHEF_DIR"); v != "" {
		return v, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".chef"), nil
}
