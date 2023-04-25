package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/CameronGorrie/sc"
	"github.com/CameronGorrie/ugens"
	"github.com/pkg/errors"
)

// Client models a connection to a running scsynth server.
type Client struct {
	server *sc.Client
	synths map[string]*sc.Synthdef
}

// NewClient creates a new sc client.
func NewClient(local, scsynth string) (*Client, error) {
	c, err := sc.NewClient("udp", local, scsynth, 5*time.Second)
	if err != nil {
		return nil, err
	}

	if _, err := c.AddDefaultGroup(); err != nil {
		return nil, fmt.Errorf("adding default group")
	}

	return &Client{
		server: c,
		synths: map[string]*sc.Synthdef{},
	}, nil
}

// FreeAll frees all nodes from scsynth.
func (c *Client) FreeAll(ctx context.Context, gids ...int) error {
	if len(gids) > 0 {
		for _, gid := range gids {
			if err := c.server.FreeAll(int32(gid)); err != nil {
				return err
			}
		}
	}

	return c.server.FreeAll(sc.DefaultGroupID)
}

// FreeGroup frees a group from scsynth.
func (c *Client) FreeGroup(ctx context.Context, gid int) error {
	return c.server.FreeAll(int32(gid))
}

// FreeNode frees a node from scsynth.
func (c *Client) FreeNode(ctx context.Context, id int) error {
	return c.server.NodeFree(int32(id))
}

// Send sends a synthdef to scsynth.
func (c *Client) Send(ctx context.Context, name string) error {
	if f, ok := ugens.Lib[name]; !ok {
		errMsg := fmt.Sprintf("no matching ugen found for name %s ", name)
		return errors.New(errMsg)
	} else {
		def := sc.NewSynthdef(name, f)
		if err := c.server.SendDef(def); err != nil {
			return err
		}

		c.synths[name] = def
	}

	return nil
}

// Play plays a synthdef.
func (c *Client) Play(ctx context.Context, name string, params []string) error {
	def := c.synths[name]
	if def == nil {
		return fmt.Errorf("no synthdef found for name %s ", name)
	}

	var err error
	ctls := map[string]float32{}
	if len(params) > 0 {
		for _, param := range params {
			a := strings.Split(param, "=")
			if len(a) < 2 {
				err = errors.Errorf("could not parse key=value from " + param)
			}

			fv, err := strconv.ParseFloat(a[1], 32)
			if err != nil {
				errors.Wrap(err, "parsing control value")
			}

			ctls[a[0]] = float32(fv)
		}
	}

	if err != nil {
		return err
	}

	if _, err := c.server.Synth(def.Name, c.server.NextSynthID(), sc.AddToTail, sc.DefaultGroupID, ctls); err != nil {
		return err
	}

	return nil
}
