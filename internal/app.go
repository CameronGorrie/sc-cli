package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/CameronGorrie/sc"
)

type App struct {
	Commands map[string]Command
}

type Command interface {
	Run(args []string) error
}

// NewApp creates the scc app and returns its exit status.
func NewApp(args []string) int {
	var app App
	if err := app.createCommands(); err != nil {
		fmt.Fprintf(os.Stderr, "[createCommands] %s", err.Error())
		return 2
	}

	cmd, err := app.parse(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[parse] %s", err.Error())
		return 2
	}

	if err := cmd.Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "[command] %s", err.Error())
		return 1
	}

	return 0
}

// createCommands connects to a running SuperCollider client and creates a map of available commands.
func (app *App) createCommands() error {
	c, err := sc.DefaultClient()
	if err != nil {
		return err
	}

	app.Commands = map[string]Command{
		"free": &Free{client: c},
		"send": &Send{client: c},
		"help": &Help{Commands: app.Commands},
	}

	return nil
}

// parse reads the command arguments provided to the app.
func (app *App) parse(args []string) (Command, error) {
	if len(args) == 0 {
		return nil, errors.New("command not provided ")
	}

	cmd, ok := app.Commands[args[0]]
	if !ok {
		return nil, errors.New("unrecognized command: " + args[0] + " ")
	}

	return cmd, nil
}
