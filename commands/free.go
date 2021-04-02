package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/CameronGorrie/sc"
)

type Free struct {
	client *sc.Client
}

func (free *Free) Run(args []string) error {
	if len(args) == 0 {
		return errors.New("No arguments provided to free")
	}

	nodeId := flag.Int("node-id", 0, "node id")
	groupId := flag.Int("group-id", 0, "group id")
	flag.Parse()

	switch {
	case *groupId != 0:
		if err := free.client.FreeAll(int32(*groupId)); err != nil {
			return err
		}
	case *nodeId != 0:
		if err := free.client.NodeFree(int32(*nodeId)); err != nil {
			return err
		}
	}

	return nil
}

func (free *Free) Usage() {
	fmt.Fprintf(os.Stderr, `
	scc free [COMMAND OPTIONS]
	COMMAND
		free					Free nodes.
	COMMAND OPTIONS
		--group-id		Free all nodes in a group.
		--node-id			Free a specific node.
	For all help options use "scc help"
	`)
}
