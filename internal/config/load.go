package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mythosmystery/chef/internal/workspace"
)

// LoadOptions configures config loading.
type LoadOptions struct {
	WorkDir       string
	Flags         FlagOverrides
	RequireGlobal bool
}

// LoadResult is the outcome of loading configuration.
type LoadResult struct {
	Config      *Config
	ProjectRoot string
	NoContext   bool
}

// Load reads global and project config, merges them, resolves paths, validates,
// and applies CLI flag overrides.
func Load(opts LoadOptions) (*LoadResult, error) {
	workDir := opts.WorkDir
	if workDir == "" {
		var err error
		workDir, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}

	projectRoot, err := workspace.FindGitRoot(workDir)
	if err != nil {
		return nil, fmt.Errorf("config: find project root: %w", err)
	}

	cfg := Defaults()

	globalPath, err := GlobalConfigPath()
	if err != nil {
		return nil, err
	}
	if opts.RequireGlobal {
		if _, err := os.Stat(globalPath); err != nil {
			if os.IsNotExist(err) {
				return nil, ErrGlobalConfigMissing{Path: globalPath}
			}
			return nil, fmt.Errorf("config: %s: %w", globalPath, err)
		}
	}
	if err := mergeFile(&cfg, globalPath); err != nil {
		return nil, fmt.Errorf("config: %s: %w", globalPath, err)
	}

	projectPath := ProjectConfigPath(projectRoot)
	if err := mergeFile(&cfg, projectPath); err != nil {
		return nil, fmt.Errorf("config: %s: %w", projectPath, err)
	}

	if err := cfg.Resolve(projectRoot); err != nil {
		return nil, fmt.Errorf("config: resolve paths: %w", err)
	}
	if err := Validate(&cfg); err != nil {
		return nil, err
	}
	if err := ApplyFlags(&cfg, opts.Flags); err != nil {
		return nil, err
	}
	if err := Validate(&cfg); err != nil {
		return nil, err
	}

	return &LoadResult{
		Config:      &cfg,
		ProjectRoot: projectRoot,
		NoContext:   NoContextFromFlags(opts.Flags),
	}, nil
}

func mergeFile(cfg *Config, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	var overlay fileConfig
	if err := json.Unmarshal(data, &overlay); err != nil {
		return err
	}
	*cfg = Merge(*cfg, overlay)
	return nil
}

