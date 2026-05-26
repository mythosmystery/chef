// Package app wires chef's subsystems and runs the program.
package app

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mythosmystery/chef/internal/tui"
)

// Flags holds CLI flags for chef.
type Flags struct {
	Init      bool
	Continue  bool
	Resume    bool
	Session   string
	NoSession bool
	Plan      string
	Model     string
	Provider  string
	Thinking  string
	Tools     string
	NoTools   bool
	NoContext bool
	Verbose   bool
	Version   bool
	Help      bool
	Message   []string
}

// Run is the program entrypoint after CLI parsing.
func Run(progName string, flags Flags) error {
	if flags.Help {
		printHelp(progName)
		return nil
	}
	if flags.Version {
		fmt.Println("chef dev")
		return nil
	}

	_ = context.Background()
	_ = flags

	deps, err := boot()
	if err != nil {
		return err
	}
	_ = deps

	p := tea.NewProgram(tui.New(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func printHelp(progName string) {
	fmt.Printf("chef — AI coding agent TUI\n\n")
	fmt.Printf("Usage: %s [options] [message...]\n\n", progName)
	fmt.Println("Options:")
	fmt.Println("  --init           Initialize project context")
	fmt.Println("  -c, --continue   Continue most recent session")
	fmt.Println("  -r, --resume     Browse and select session")
	fmt.Println("  --session PATH   Specific session file")
	fmt.Println("  --no-session     Ephemeral mode")
	fmt.Println("  --plan PROMPT    Start in plan mode")
	fmt.Println("  --model NAME     Override model")
	fmt.Println("  --provider NAME  Override provider")
	fmt.Println("  --thinking LEVEL Thinking level: off, low, medium, high")
	fmt.Println("  -t, --tools LIST Comma-separated tool allowlist")
	fmt.Println("  --no-tools       Disable all tools")
	fmt.Println("  --no-context     Skip project context injection")
	fmt.Println("  --verbose        Verbose logging")
	fmt.Println("  -v, --version    Print version")
	fmt.Println("  -h, --help       Show help")
}

// ExitOnError prints err to stderr and exits with code 1.
func ExitOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
