package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/mythosmystery/chef/internal/config"
	"github.com/mythosmystery/chef/internal/workspace"
)

type configAnswers struct {
	Provider         string
	CustomBaseURL    string
	CustomAPIKeyEnv  string
	Model            string
	Thinking         string
	Theme            string
	LightEnabled     bool
	LightProvider    string
	LightCustomURL   string
	LightCustomKey   string
	LightModel       string
}

func runConfig(progName string, args []string) error {
	fs := flag.NewFlagSet("config", flag.ContinueOnError)
	project := fs.Bool("project", false, "write project config instead of global")
	help := fs.Bool("h", false, "show help")
	fs.BoolVar(help, "help", false, "show help")
	fs.SetOutput(os.Stderr)

	if err := fs.Parse(args); err != nil {
		return err
	}
	if *help {
		printConfigHelp(progName)
		return nil
	}

	base, targetPath, err := configTarget(*project)
	if err != nil {
		return err
	}

	answers := answersFromConfig(base)
	if err := runConfigWizard(&answers); err != nil {
		return err
	}

	cfg := buildConfigFromAnswers(base, answers)
	if err := config.Validate(&cfg); err != nil {
		return err
	}
	if err := config.WriteFile(targetPath, cfg); err != nil {
		return err
	}

	fmt.Printf("Wrote config to %s\n", targetPath)
	printAPIKeyReminder(&cfg)
	return nil
}

func printConfigHelp(progName string) {
	fmt.Printf("Usage: %s config [--project]\n\n", progName)
	fmt.Println("Interactive setup wizard for chef configuration.")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --project    Write to .chef/config.json in the current git repo")
	fmt.Println("  -h, --help   Show this help")
}

func configTarget(project bool) (config.Config, string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return config.Config{}, "", err
	}

	if project {
		projectRoot, err := workspace.FindGitRoot(cwd)
		if err != nil {
			return config.Config{}, "", err
		}
		gitPath := filepath.Join(projectRoot, ".git")
		if _, err := os.Stat(gitPath); err != nil {
			if os.IsNotExist(err) {
				return config.Config{}, "", fmt.Errorf("not in a git repository (--project requires a git repo)")
			}
			return config.Config{}, "", err
		}

		result, err := config.Load(config.LoadOptions{WorkDir: cwd})
		if err != nil {
			return config.Config{}, "", err
		}
		return *result.Config, config.ProjectConfigPath(projectRoot), nil
	}

	return configTargetGlobal()
}

func configTargetGlobal() (config.Config, string, error) {
	base := config.Defaults()
	path, err := config.GlobalConfigPath()
	if err != nil {
		return config.Config{}, "", err
	}
	return base, path, nil
}

func answersFromConfig(cfg config.Config) configAnswers {
	a := configAnswers{
		Provider: cfg.Provider,
		Model:    cfg.Model,
		Thinking: cfg.Thinking,
		Theme:    cfg.Theme,
	}
	if cfg.Light != nil {
		a.LightEnabled = true
		a.LightProvider = cfg.Light.Provider
		a.LightModel = cfg.Light.Model
	}
	if pc, ok := cfg.Providers["custom"]; ok {
		a.CustomBaseURL = pc.BaseURL
		a.CustomAPIKeyEnv = pc.APIKeyEnv
	}
	return a
}

func buildConfigFromAnswers(base config.Config, a configAnswers) config.Config {
	cfg := base
	cfg.Provider = a.Provider
	cfg.Model = a.Model
	cfg.Thinking = a.Thinking
	cfg.Theme = a.Theme
	if a.LightEnabled {
		cfg.Light = &config.LightConfig{
			Provider: a.LightProvider,
			Model:    a.LightModel,
		}
	} else {
		cfg.Light = nil
	}
	applyCustomProviderConfig(&cfg, a.Provider, a.CustomBaseURL, a.CustomAPIKeyEnv)
	if a.LightEnabled && a.LightProvider == "custom" {
		applyCustomProviderConfig(&cfg, "custom", a.LightCustomURL, a.LightCustomKey)
	}
	return cfg
}

func applyCustomProviderConfig(cfg *config.Config, providerName, baseURL, apiKeyEnv string) {
	if providerName != "custom" {
		return
	}
	if cfg.Providers == nil {
		cfg.Providers = make(map[string]config.ProviderConfig)
	}
	pc := cfg.Providers["custom"]
	pc.BaseURL = baseURL
	if apiKeyEnv != "" {
		pc.APIKeyEnv = apiKeyEnv
	} else if pc.APIKeyEnv == "" {
		pc.APIKeyEnv = "CUSTOM_API_KEY"
	}
	cfg.Providers["custom"] = pc
}

func runConfigWizard(a *configAnswers) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Provider").
				Description("LLM provider for the main agent").
				Options(providerSelectOptions()...).
				Value(&a.Provider),
			huh.NewInput().
				Title("Model").
				Description("Main model name").
				Placeholder("gpt-4o, claude-sonnet-4-20250514").
				Value(&a.Model).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("model is required")
					}
					return nil
				}),
			huh.NewSelect[string]().
				Title("Thinking").
				Description("Extended reasoning level").
				Options(
					huh.NewOption("Off", "off"),
					huh.NewOption("Low", "low"),
					huh.NewOption("Medium", "medium"),
					huh.NewOption("High", "high"),
				).
				Value(&a.Thinking),
			huh.NewSelect[string]().
				Title("Theme").
				Description("TUI color theme").
				Options(
					huh.NewOption("Dark", "dark"),
					huh.NewOption("Light", "light"),
				).
				Value(&a.Theme),
			huh.NewConfirm().
				Title("Light model").
				Description("Configure a separate model for mini-agents?").
				Value(&a.LightEnabled),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	if a.Provider == "custom" {
		if a.CustomAPIKeyEnv == "" {
			a.CustomAPIKeyEnv = "CUSTOM_API_KEY"
		}
		customForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Base URL").
					Description("OpenAI-compatible API base URL (e.g. https://host/v1)").
					Placeholder("https://api.example.com/v1").
					Value(&a.CustomBaseURL).
					Validate(func(s string) error {
						if s == "" {
							return fmt.Errorf("base URL is required for custom provider")
						}
						return nil
					}),
				huh.NewInput().
					Title("API key env var").
					Description("Environment variable name holding the API key").
					Placeholder("CUSTOM_API_KEY").
					Value(&a.CustomAPIKeyEnv).
					Validate(func(s string) error {
						if s == "" {
							return fmt.Errorf("API key env var name is required")
						}
						return nil
					}),
			),
		)
		if err := customForm.Run(); err != nil {
			return err
		}
	}

	if a.LightEnabled {
		if a.LightProvider == "" {
			a.LightProvider = a.Provider
		}
		lightForm := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Light provider").
					Description("Provider for mini-agents").
					Options(providerSelectOptions()...).
					Value(&a.LightProvider),
				huh.NewInput().
					Title("Light model").
					Description("Model for mini-agents").
					Placeholder("gpt-4o-mini, claude-haiku-4-20250414").
					Value(&a.LightModel).
					Validate(func(s string) error {
						if s == "" {
							return fmt.Errorf("light model is required")
						}
						return nil
					}),
			),
		)
		if err := lightForm.Run(); err != nil {
			return err
		}

		if a.LightProvider == "custom" {
			if a.LightCustomKey == "" {
				a.LightCustomKey = "CUSTOM_API_KEY"
			}
			if a.LightCustomURL == "" {
				a.LightCustomURL = a.CustomBaseURL
			}
			lightCustomForm := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Light base URL").
						Description("OpenAI-compatible API base URL for mini-agents").
						Placeholder("https://api.example.com/v1").
						Value(&a.LightCustomURL).
						Validate(func(s string) error {
							if s == "" {
								return fmt.Errorf("base URL is required for custom provider")
							}
							return nil
						}),
					huh.NewInput().
						Title("Light API key env var").
						Value(&a.LightCustomKey).
						Validate(func(s string) error {
							if s == "" {
								return fmt.Errorf("API key env var name is required")
							}
							return nil
						}),
				),
			)
			if err := lightCustomForm.Run(); err != nil {
				return err
			}
		}
	}

	return nil
}

func providerSelectOptions() []huh.Option[string] {
	return []huh.Option[string]{
		huh.NewOption("OpenAI", "openai"),
		huh.NewOption("Anthropic", "anthropic"),
		huh.NewOption("Custom (OpenAI-compatible)", "custom"),
	}
}

func printAPIKeyReminder(cfg *config.Config) {
	envVars := collectAPIKeyEnvs(cfg)
	if len(envVars) == 0 {
		return
	}
	fmt.Println()
	fmt.Println("Set your API key(s) in the environment:")
	for _, env := range envVars {
		fmt.Printf("  export %s=...\n", env)
	}
}

func collectAPIKeyEnvs(cfg *config.Config) []string {
	seen := make(map[string]struct{})
	var out []string
	add := func(name string, pc config.ProviderConfig) {
		env := pc.APIKeyEnv
		if env == "" {
			switch name {
			case "openai":
				env = "OPENAI_API_KEY"
			case "anthropic":
				env = "ANTHROPIC_API_KEY"
			default:
				return
			}
		}
		if _, ok := seen[env]; ok {
			return
		}
		seen[env] = struct{}{}
		out = append(out, env)
	}
	mainName, _ := cfg.MainModel()
	add(mainName, cfg.ActiveProviderConfig())
	lightName, _ := cfg.LightModel()
	if lightName != mainName {
		add(lightName, cfg.LightProviderConfig())
	}
	return out
}
