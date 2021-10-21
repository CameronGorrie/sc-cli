package commands

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/CameronGorrie/sc"
)

type App struct {
	scsynthAddr string
	Commands    map[string]Command
}

type Command interface {
	Run(args []string) error
}

// NewApp creates the sc-cli app and returns its exit status.
func NewApp(args []string) int {
	var app App
	err := app.createCommands()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] %s", err.Error())
		return 2
	}

	cmd, err := app.parseCommandFromArgs(args)
	if err := cmd.Run(args); err != nil {
		return 1
	}

	return 0
}

func (app *App) createCommands() error {
	scc, err := sc.NewClient("udp", "127.0.0.1:0", sc.DefaultScsynthAddr, 5*time.Second)
	if err != nil {
		return err
	}

	app.Commands = map[string]Command{
		"free": &Free{client: scc},
	}
	app.Commands["help"] = &Help{Commands: app.Commands}

	return nil
}

func (app *App) parseCommandFromArgs(args []string) (Command, error) {
	if len(args) == 0 {
		return nil, errors.New("command not provided")
	}

	cmd, ok := app.Commands[args[0]]
	if !ok {
		return nil, errors.New("unrecognized command: " + args[0])
	}

	return cmd, nil
}
