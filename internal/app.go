package cmd

import (
	"errors"
	"fmt"
	"os"
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

// createCommands creates a map of available commands.
func (app *App) createCommands() error {
	app.Commands = map[string]Command{
		"free": &Free{},
		"send": &Send{},
		"help": &Help{Commands: app.Commands},
	}

	return nil
}

// parse reads the command arguments provided to the app.
func (app *App) parse(args []string) (Command, error) {
	if len(args) == 0 {
		return nil, errors.New("command not provided ")
	}

	if err := app.createCommands(); err != nil {
		return nil, err
	}

	cmd, ok := app.Commands[args[0]]
	if !ok {
		return nil, errors.New("unrecognized command: " + args[0] + " ")
	}

	return cmd, nil
}
