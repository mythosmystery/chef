package app

import (
	"github.com/mythosmystery/chef/internal/bus"
	"github.com/mythosmystery/chef/internal/config"
)

// Dependencies holds wired subsystems for a chef run.
type Dependencies struct {
	Config *config.Config
	Bus    *bus.Bus
}

// boot constructs config, provider, tools, agent, and TUI dependencies.
func boot() (*Dependencies, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Config: cfg,
		Bus:    bus.New(),
	}, nil
}
