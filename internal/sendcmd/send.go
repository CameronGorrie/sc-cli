package sendcmd

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"

	"github.com/CameronGorrie/scc/internal/rootcmd"
	"github.com/peterbourgon/ff/v3/ffcli"
)

// Config for the send subcommand, including a reference
// to the global config, for access to global flags.
type Config struct {
	rootConfig *rootcmd.Config
	out        io.Writer
}

// New creates a new ffcli.Command for the free subcommand.
func New(rootConfig *rootcmd.Config, out io.Writer) *ffcli.Command {
	cfg := Config{
		rootConfig: rootConfig,
		out:        out,
	}

	fs := flag.NewFlagSet("send", flag.ExitOnError)

	rootConfig.RegisterFlags(fs)

	return &ffcli.Command{
		Name:       "send",
		ShortUsage: "scc send <ugen1> <ugen2> ...",
		ShortHelp:  "Send Ugen functions to a running instance of scsynth",
		FlagSet:    fs,
		Exec:       cfg.Exec,
	}
}

// Exec function for this command.
func (c *Config) Exec(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return errors.New("send requires at least 1 argument")
	}

	for _, name := range args {
		if err := c.rootConfig.Client.Send(ctx, name); err != nil {
			return err
		}

		if c.rootConfig.Verbose {
			fmt.Fprintf(c.out, "%s ugen sent to server\n", name)
		}
	}

	return nil
}
