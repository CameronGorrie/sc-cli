package freecmd

import (
	"context"
	"flag"
	"io"

	"github.com/CameronGorrie/scc/internal/rootcmd"
	"github.com/peterbourgon/ff/v3/ffcli"
)

// Config for the free subcommand, including a reference
// to the global config, for access to global flags.
type Config struct {
	rootConfig *rootcmd.Config
	out        io.Writer
	all        bool
	gid        int
	id         int
}

// New creates a new ffcli.Command for the free subcommand.
func New(rootConfig *rootcmd.Config, out io.Writer) *ffcli.Command {
	cfg := Config{
		rootConfig: rootConfig,
		out:        out,
	}

	fs := flag.NewFlagSet("scc free", flag.ExitOnError)
	fs.IntVar(&cfg.gid, "gid", 0, "group id")
	fs.IntVar(&cfg.id, "id", 0, "node id")
	fs.BoolVar(&cfg.all, "a", false, "all")

	rootConfig.RegisterFlags(fs)

	return &ffcli.Command{
		Name:       "free",
		ShortUsage: "scc free [flags]",
		ShortHelp:  "Free nodes from scsynth",
		FlagSet:    fs,
		Exec:       cfg.Exec,
	}
}

// Exec function for this command.
func (c *Config) Exec(ctx context.Context, args []string) error {
	if c.all {
		return c.rootConfig.Client.FreeAll(ctx)
	}

	if c.gid != 0 {
		return c.rootConfig.Client.FreeGroup(ctx, c.gid)
	}

	if c.id != 0 {
		return c.rootConfig.Client.FreeNode(ctx, c.id)
	}

	return nil
}
