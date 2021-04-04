package commands

import (
	"errors"
	"fmt"
	"os"
)

type Help struct {
	Commands map[string]Command
}

func (help *Help) Run(args []string) error {
	if len(args) == 0 {
		helpMessage()
		return nil
	}

	cmd, ok := help.Commands[args[0]]
	if !ok {
		return errors.New("unrecognized command: " + args[0])
	}
	cmd.Usage()
	return nil
}

func (help *Help) Usage() {
	helpMessage()
}

func helpMessage() {
	fmt.Fprintf(os.Stderr, `
sc-cli [GLOBAL OPTIONS] COMMAND [COMMAND OPTIONS]
GLOBAL OPTIONS
	-h			Remote scsynth host.
	-p			Remote scsynth port.
COMMANDS
	free		Free nodes.
For help with a particular command, "scc help COMMAND"
	`)
}
