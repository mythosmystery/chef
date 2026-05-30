package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mythosmystery/chef/internal/app"
	"github.com/mythosmystery/chef/internal/cli"
	"github.com/mythosmystery/chef/internal/config"
)

func main() {
	if len(os.Args) > 1 && !strings.HasPrefix(os.Args[1], "-") {
		os.Exit(cli.Run(os.Args[0], os.Args[1:]))
	}

	flags, err := ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := app.Run(os.Args[0], flags); err != nil {
		var missing config.ErrGlobalConfigMissing
		if errors.As(err, &missing) {
			fmt.Fprint(os.Stderr, err.Error())
		} else {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		os.Exit(1)
	}
}
