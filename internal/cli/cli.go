// Package cli implements chef subcommands.
package cli

import (
	"fmt"
	"os"
)

type handler func(progName string, args []string) error

var commands = map[string]handler{
	"config": runConfig,
}

// Run dispatches a subcommand and returns an exit code.
func Run(progName string, args []string) int {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: %s <command>\n\nCommands:\n", progName)
		for name := range commands {
			fmt.Fprintf(os.Stderr, "  %s\n", name)
		}
		return 1
	}

	name := args[0]
	fn, ok := commands[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown command %q\n\nCommands:\n", name)
		for n := range commands {
			fmt.Fprintf(os.Stderr, "  %s\n", n)
		}
		return 1
	}

	if err := fn(progName, args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	return 0
}
