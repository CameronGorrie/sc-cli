package cmd

import (
	"fmt"
	"time"

	"github.com/CameronGorrie/sc"
)

func NewClient(scsynthAddr string) (*sc.Client, error) {
	c, err := sc.NewClient("udp", "0.0.0.0:0", scsynthAddr, 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("creating sc client")
	}

	if _, err := c.AddDefaultGroup(); err != nil {
		return nil, fmt.Errorf("adding default group")
	}

	return c, nil
}
