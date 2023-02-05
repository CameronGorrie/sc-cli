package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/CameronGorrie/sc"
)

type App struct {
	commands    map[string]Command
	scsynthAddr string
	help        bool
}

type Command struct {
	fs  *flag.FlagSet
	cmd Action
}

type Action interface {
	Run(c *sc.Client) error
	ParseFlags(fs *flag.FlagSet, args []string) error
}

// NewApp creates the scc app and returns its exit status.
func NewApp(args []string) int {
	var app App

	app.createCommands()
	app.setupCommonFlags()

	sCmd := app.commands[args[0]]
	if sCmd.fs == nil {
		fmt.Fprintf(os.Stderr, "[parse] %s\n", args[1])
		return 2
	}

	if err := sCmd.cmd.ParseFlags(sCmd.fs, args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "[] %s\n", err.Error())
		return 2
	}

	c, err := NewClient(app.scsynthAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[client] %s\n", err.Error())
		return 1
	}

	if printUsage(args) {
		sCmd.fs.Usage()
		return 0
	}

	if err := sCmd.cmd.Run(c); err != nil {
		fmt.Fprintf(os.Stderr, "[command] %s\n", err.Error())
		return 1
	}

	return 0
}

// createCommands creates a map of available commands with their flagsets.
func (app *App) createCommands() {
	freeCmd := flag.NewFlagSet("free", flag.ContinueOnError)
	sendCmd := flag.NewFlagSet("send", flag.ContinueOnError)

	app.commands = map[string]Command{
		freeCmd.Name(): {fs: freeCmd, cmd: &Free{}},
		sendCmd.Name(): {fs: sendCmd, cmd: &Send{}},
	}
}

// setupCommonFlags creates a common flag for each command.
func (app *App) setupCommonFlags() {
	for _, sCmd := range app.commands {
		sCmd.fs.StringVar(
			&app.scsynthAddr,
			"u",
			sc.DefaultScsynthAddr,
			"remote address for scsynth",
		)

		sCmd.fs.BoolVar(
			&app.help,
			"h",
			false,
			"print command usage",
		)
	}
}

func printUsage(args []string) bool {
	fmt.Println(args)
	for _, arg := range args {
		if arg == "-h" {
			return true
		}
	}

	return false
}
