package cmd

import (
	"errors"
	"flag"
)

type Free struct {
	port    int
	freeAll bool
	groupId int
	nodeId  int
}

func (f *Free) Run(args []string) error {
	if len(args[1:]) == 0 {
		return errors.New("no arguments provided to free ")
	}

	fs := flag.NewFlagSet("free", flag.ContinueOnError)
	fs.IntVar(&f.groupId, "gid", 0, "group id")
	fs.IntVar(&f.nodeId, "id", 0, "node id")
	fs.IntVar(&f.port, "u", 57120, "UDP port")
	fs.BoolVar(&f.freeAll, "a", true, "all")

	if err := fs.Parse(args[1:]); err != nil {
		return err
	}

	c, err := NewClient(f.port)
	if err != nil {
		return err
	}

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
	}

	if f.nodeId != 0 {
		if err := c.NodeFree(int32(f.nodeId)); err != nil {
			return err
		}
	}

	return nil
}
