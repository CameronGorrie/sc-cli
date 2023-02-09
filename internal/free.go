package cmd

import (
	"errors"
	"flag"

	"github.com/CameronGorrie/sc"
)

type Free struct {
	freeAll bool
	groupId int
	nodeId  int
}

func (f *Free) Run(c *sc.Client) error {
	if f.freeAll {
		if err := c.FreeAll(int32(f.groupId)); err != nil {
			return err
		}

		return nil
	}

	if f.groupId != 0 {
		if err := c.FreeAll(int32(f.groupId)); err != nil {
			return err
		}

		return nil
	}

	if f.nodeId != 0 {
		if err := c.NodeFree(int32(f.nodeId)); err != nil {
			return err
		}
	}

	return nil
}

func (f *Free) ParseFlags(fs *flag.FlagSet, args []string) error {
	if len(args) == 0 {
		return errors.New("no arguments provided to free ")
	}

	fs.IntVar(&f.groupId, "gid", int(sc.DefaultGroupID), "group id")
	fs.IntVar(&f.nodeId, "id", 0, "node id")
	fs.BoolVar(&f.freeAll, "a", true, "all")

	if err := fs.Parse(args[0:]); err != nil {
		return err
	}

	return nil
}
