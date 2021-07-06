package commands

import (
	"errors"
	"flag"

	"github.com/CameronGorrie/sc"
)

type Free struct {
	groupId int
	nodeId  int
	client  *sc.Client
}

func (free *Free) Run(args []string) error {
	if len(args) == 0 {
		return errors.New("No arguments provided to free")
	}

	fs := flag.NewFlagSet("free", flag.ContinueOnError)
	fs.IntVar(&free.groupId, "group-id", 0, "group id")
	fs.IntVar(&free.nodeId, "node-id", 0, "node id")

	if err := fs.Parse(args[1:]); err != nil {
		return err
	}

	for i := 0; i < len(args); i++ {
		switch {
		case free.groupId != 0:
			if err := free.client.FreeAll(int32(free.groupId)); err != nil {
				return err
			}
		case free.nodeId != 0:
			if err := free.client.NodeFree(int32(free.nodeId)); err != nil {
				return err
			}
		}
	}

	return nil
}
