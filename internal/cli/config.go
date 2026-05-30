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
	Provider      string
	Model         string
	Thinking      string
	Theme         string
	LightEnabled  bool
	LightProvider string
	LightModel    string
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
	printAPIKeyReminder(answers.Provider, answers.LightEnabled, answers.LightProvider)
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
	return cfg
}

func runConfigWizard(a *configAnswers) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Provider").
				Description("LLM provider for the main agent").
				Options(
					huh.NewOption("OpenAI", "openai"),
					huh.NewOption("Anthropic", "anthropic"),
				).
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

	if a.LightEnabled {
		if a.LightProvider == "" {
			a.LightProvider = a.Provider
		}
		lightForm := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Light provider").
					Description("Provider for mini-agents").
					Options(
						huh.NewOption("OpenAI", "openai"),
						huh.NewOption("Anthropic", "anthropic"),
					).
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
	}

	return nil
}

func printAPIKeyReminder(provider string, lightEnabled bool, lightProvider string) {
	needsOpenAI := provider == "openai" || (lightEnabled && lightProvider == "openai")
	needsAnthropic := provider == "anthropic" || (lightEnabled && lightProvider == "anthropic")

	if !needsOpenAI && !needsAnthropic {
		return
	}

	fmt.Println()
	fmt.Println("Set your API key(s) in the environment:")
	if needsOpenAI {
		fmt.Println("  export OPENAI_API_KEY=...")
	}
	if needsAnthropic {
		fmt.Println("  export ANTHROPIC_API_KEY=...")
	}
}
