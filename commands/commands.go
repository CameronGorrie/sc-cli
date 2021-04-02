package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/CameronGorrie/sc"
)

type Command interface {
	Run(args []string) error
	Usage() string
}

type appEnv struct {
	Commands map[string]Command
}

// CLI runs the SuperCollider command line app and returns its exit status.
func CLI(args []string) int {
	var app appEnv
	cmd, err := app.getCommandFromArgs(args)

	if err != nil {
		fmt.Fprintf(os.Stderr, `There was a problem`, err.Error())
		return 2
	}

	if err := cmd.Run(args[1:]); err != nil {
		usg := cmd.Usage()
		fmt.Fprintf(os.Stderr, usg, err)
		return 1
	}
	return 0
}

func (app *appEnv) getCommandFromArgs(args []string) (Command, error) {
	if len(args) == 0 {
		return nil, errors.New("Command not found")
	}

	// use scsynth address from config flag or provide default
	var (
		port int
		host string
	)
	flag.IntVar(&port, "p", 57120, "scsynth port")
	flag.StringVar(&host, "h", "127.0.0.1", "scsynth host")
	serverAddr := fmt.Sprintf("%s:%d", host, port)

	// create new sc client
	scc, err := sc.NewClient("udp", "127.0.0.1:0", serverAddr, 5*time.Second)
	if err != nil {
		return nil, err
	}

	// create the commands map
	app.Commands = map[string]Command{
		"free": &Free{client: scc},
	}

	// get command from first arg `app.Commands[args[0]]`
	cmd, ok := app.Commands[args[0]]
	if !ok {
		return nil, errors.New("There was a problem")
	}

	return cmd, nil
}
