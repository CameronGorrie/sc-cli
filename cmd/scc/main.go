package main

import (
	"context"
	"fmt"
	"os"
	"time"

	client "github.com/CameronGorrie/scc"
	"github.com/CameronGorrie/scc/internal/freecmd"
	"github.com/CameronGorrie/scc/internal/playcmd"
	"github.com/CameronGorrie/scc/internal/rootcmd"
	"github.com/CameronGorrie/scc/internal/sendcmd"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	var (
		out         = os.Stdout
		cmd, cfg    = rootcmd.New()
		freeCommand = freecmd.New(cfg, out)
		sendCommand = sendcmd.New(cfg, out)
		playCommand = playcmd.New(cfg, out)
	)

	cmd.Subcommands = []*ffcli.Command{
		freeCommand,
		sendCommand,
		playCommand,
	}

	if err := cmd.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error during Parse: %v\n", err)
		os.Exit(1)
	}

	client, err := client.NewClient(cfg.LocalAddr, cfg.SynthAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error constructing SuperCollider client: %v\n", err)
		os.Exit(1)
	}

	cfg.Client = client
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
	defer cancel()

	if err := cmd.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
