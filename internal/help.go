package cmd

import (
	"flag"
	"fmt"
	"os"
)

type Help struct {
	Subcommands map[string]*flag.FlagSet
}

func (help *Help) Run(args []string) error {
	Usage()

	return nil
}

func Usage() {
	fmt.Fprintf(os.Stderr, `
scc COMMAND [COMMAND OPTIONS]
COMMANDS
	free		Free groups and nodes.
For help with a particular command, "scc COMMAND -h"
	`)
}
