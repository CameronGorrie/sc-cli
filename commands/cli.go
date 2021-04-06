package commands

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/CameronGorrie/sc"
)

type Command interface {
	Run(args []string) error
}

type appEnv struct {
	scsynthAddr string
	Commands    map[string]Command
}

// CLI runs the SuperCollider command line app and returns its exit status.
func CLI(args []string) int {
	var app appEnv
	cmd, err := app.getCommandFromArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] %s", err.Error())
		return 2
	}
	if err := cmd.Run(args); err != nil {
		return 1
	}
	return 0
}

func (app *appEnv) getCommandFromArgs(args []string) (Command, error) {
	if len(args) == 0 {
		return nil, errors.New("command not provided")
	}

	scc, err := sc.NewClient("udp", "127.0.0.1:0", sc.DefaultScsynthAddr, 5*time.Second)
	if err != nil {
		return nil, err
	}

	app.Commands = map[string]Command{
		"free": &Free{client: scc},
	}
	app.Commands["help"] = &Help{Commands: app.Commands}

	cmd, ok := app.Commands[args[0]]
	if !ok {
		return nil, errors.New("unrecognized command: " + args[0])
	}

	return cmd, nil
}
