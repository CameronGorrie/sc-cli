package cmd

import (
	"fmt"
	"time"

	"github.com/CameronGorrie/sc"
)

func NewClient(port int) (*sc.Client, error) {
	c, err := sc.NewClient("udp", "0.0.0.0:0", fmt.Sprintf("0.0.0.0:%d", port), 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("creating sc client")
	}

	if _, err := c.AddDefaultGroup(); err != nil {
		return nil, fmt.Errorf("adding default group")
	}

	return c, nil
}
