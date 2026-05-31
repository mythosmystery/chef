package app

import (
	"os"

	"github.com/mythosmystery/chef/internal/bus"
	"github.com/mythosmystery/chef/internal/config"
	"github.com/mythosmystery/chef/pkg/provider"
)

// Dependencies holds wired subsystems for a chef run.
type Dependencies struct {
	Config        *config.Config
	ProjectRoot   string
	NoContext     bool
	Bus           *bus.Bus
	Provider      provider.Provider
	LightProvider provider.Provider
}

// boot constructs config, provider, tools, agent, and TUI dependencies.
func boot(flags Flags) (*Dependencies, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	result, err := config.Load(config.LoadOptions{
		WorkDir:       cwd,
		Flags:         flagOverrides(flags),
		RequireGlobal: true,
	})
	if err != nil {
		return nil, err
	}

	mainProv, lightProv, err := buildProviders(result.Config)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Config:        result.Config,
		ProjectRoot:   result.ProjectRoot,
		NoContext:     result.NoContext,
		Bus:           bus.New(),
		Provider:      mainProv,
		LightProvider: lightProv,
	}, nil
}

func flagOverrides(f Flags) config.FlagOverrides {
	return config.FlagOverrides{
		Provider:  f.Provider,
		Model:     f.Model,
		Thinking:  f.Thinking,
		Tools:     f.Tools,
		NoTools:   f.NoTools,
		NoContext: f.NoContext,
	}
}
