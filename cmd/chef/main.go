package main

import (
	"fmt"
	"os"

	"github.com/mythosmystery/chef/internal/app"
)

func main() {
	flags, err := ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := app.Run(os.Args[0], flags); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
