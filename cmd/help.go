package cmd

import (
	"fmt"
	"os"
)

type Help struct {
	Commands map[string]Command
}

func (help *Help) Run(args []string) error {
	Usage()

	return nil
}

func Usage() {
	fmt.Fprintf(os.Stderr, `
sc-cli COMMAND [COMMAND OPTIONS]
COMMANDS
	free		Free groups and nodes.
For help with a particular command, "sc-cli COMMAND -h"
	`)
}
