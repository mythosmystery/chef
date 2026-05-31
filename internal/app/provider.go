package app

import (
	"fmt"
	"os"

	"github.com/mythosmystery/chef/internal/config"
	"github.com/mythosmystery/chef/pkg/provider"
	"github.com/mythosmystery/chef/pkg/provider/anthropic"
	"github.com/mythosmystery/chef/pkg/provider/openai"
)

func buildProviders(cfg *config.Config) (main provider.Provider, light provider.Provider, err error) {
	mainName, _ := cfg.MainModel()
	main, err = buildProvider(mainName, cfg.ActiveProviderConfig())
	if err != nil {
		return nil, nil, fmt.Errorf("main provider: %w", err)
	}

	lightName, _ := cfg.LightModel()
	if lightName == mainName {
		return main, main, nil
	}

	light, err = buildProvider(lightName, cfg.LightProviderConfig())
	if err != nil {
		return nil, nil, fmt.Errorf("light provider: %w", err)
	}
	return main, light, nil
}

func buildProvider(name string, pc config.ProviderConfig) (provider.Provider, error) {
	apiKey, err := resolveAPIKey(name, pc)
	if err != nil {
		return nil, err
	}

	switch name {
	case "openai", "custom":
		return openai.New(openai.Options{
			APIKey:       apiKey,
			BaseURL:      pc.BaseURL,
			Organization: pc.Organization,
			Project:      pc.Project,
			Headers:      pc.Headers,
			Timeout:      pc.Timeout.Duration,
			MaxRetries:   pc.MaxRetries,
			Name:         name,
		}), nil
	case "anthropic":
		return anthropic.New(anthropic.Options{
			APIKey:     apiKey,
			BaseURL:    pc.BaseURL,
			APIVersion: pc.Version,
			Beta:       pc.Beta,
			Headers:    pc.Headers,
			Timeout:    pc.Timeout.Duration,
			MaxRetries: pc.MaxRetries,
		}), nil
	default:
		return nil, fmt.Errorf("unknown provider %q", name)
	}
}

func resolveAPIKey(providerName string, pc config.ProviderConfig) (string, error) {
	envVar := pc.APIKeyEnv
	if envVar == "" {
		envVar = defaultAPIKeyEnv(providerName)
	}
	if envVar == "" {
		return "", fmt.Errorf("no apiKeyEnv configured for provider %q", providerName)
	}
	key := os.Getenv(envVar)
	if key == "" {
		return "", fmt.Errorf("environment variable %s is not set (required for provider %q)", envVar, providerName)
	}
	return key, nil
}

func defaultAPIKeyEnv(providerName string) string {
	switch providerName {
	case "openai":
		return "OPENAI_API_KEY"
	case "anthropic":
		return "ANTHROPIC_API_KEY"
	default:
		return ""
	}
}
