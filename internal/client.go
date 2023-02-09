package cmd

import (
	"fmt"
	"time"

	"github.com/CameronGorrie/sc"
)

func NewClient(scsynthAddr string) (*sc.Client, error) {
	c, err := sc.NewClient("udp", "0.0.0.0:0", scsynthAddr, 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("%s creating sc client", err)
	}

	if _, err := c.AddDefaultGroup(); err != nil {
		return nil, fmt.Errorf("%s adding default group", err)
	}

	return c, nil
}
