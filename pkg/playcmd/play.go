package playcmd

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"

	"github.com/CameronGorrie/scc/pkg/rootcmd"
	"github.com/peterbourgon/ff/v3/ffcli"
)

// Config for the play subcommand, including a reference
// to the global config, for access to global flags.
type Config struct {
	rootConfig *rootcmd.Config
	out        io.Writer
	send       bool
}

// New creates a new ffcli.Command for the free subcommand.
func New(rootConfig *rootcmd.Config, out io.Writer) *ffcli.Command {
	cfg := Config{
		rootConfig: rootConfig,
		out:        out,
	}

	fs := flag.NewFlagSet("play", flag.ExitOnError)
	fs.BoolVar(&cfg.send, "send", false, "send the synthdef to scsynth")

	rootConfig.RegisterFlags(fs)

	return &ffcli.Command{
		Name:       "play",
		ShortUsage: "scc play [flags] <name> <ctls key=value...>",
		ShortHelp:  "Play a synthdef on a running instance of scsynth",
		FlagSet:    fs,
		Exec:       cfg.Exec,
	}
}

// Exec function for this command.
func (c *Config) Exec(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return errors.New("play requires at least 1 argument")
	}

	name := args[0]

	if c.send {
		if err := c.rootConfig.Client.Send(ctx, name); err != nil {
			return err
		}
	}

	if err := c.rootConfig.Client.Play(ctx, name, args[0:]); err != nil {
		return err
	}

	if c.rootConfig.Verbose {
		fmt.Fprintf(c.out, "%s playing synth\n", name)
	}

	return nil
}
