package main

import (
	"flag"

	"github.com/mythosmystery/chef/internal/app"
)

// ParseFlags parses command-line flags and returns the remaining args.
func ParseFlags(args []string) (app.Flags, error) {
	fs := flag.NewFlagSet("chef", flag.ContinueOnError)
	f := app.Flags{}

	fs.BoolVar(&f.Init, "init", false, "initialize project context")
	fs.BoolVar(&f.Continue, "c", false, "continue most recent session")
	fs.BoolVar(&f.Continue, "continue", false, "continue most recent session")
	fs.BoolVar(&f.Resume, "r", false, "browse and select session")
	fs.BoolVar(&f.Resume, "resume", false, "browse and select session")
	fs.StringVar(&f.Session, "session", "", "specific session file")
	fs.BoolVar(&f.NoSession, "no-session", false, "ephemeral mode")
	fs.StringVar(&f.Plan, "plan", "", "start in plan mode")
	fs.StringVar(&f.Model, "model", "", "override model")
	fs.StringVar(&f.Provider, "provider", "", "override provider")
	fs.StringVar(&f.Thinking, "thinking", "", "thinking level: off, low, medium, high")
	fs.StringVar(&f.Tools, "tools", "", "comma-separated tool allowlist")
	fs.BoolVar(&f.NoTools, "no-tools", false, "disable all tools")
	fs.BoolVar(&f.NoContext, "no-context", false, "skip project context injection")
	fs.BoolVar(&f.Verbose, "verbose", false, "verbose logging")
	fs.BoolVar(&f.Version, "v", false, "print version")
	fs.BoolVar(&f.Version, "version", false, "print version")
	fs.BoolVar(&f.Help, "h", false, "show help")
	fs.BoolVar(&f.Help, "help", false, "show help")

	if err := fs.Parse(args); err != nil {
		return app.Flags{}, err
	}
	f.Message = fs.Args()
	return f, nil
}
