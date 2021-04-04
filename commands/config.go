package commands

import (
	"flag"
	"fmt"
)

// ParseConfig returns the scsynth listening address
func ParseConfig() string {
	host := flag.String("h", "127.0.0.1", "host")
	port := flag.Int("p", 57120, "port")
	flag.Parse()

	return fmt.Sprintf("%s:%d", *host, *port)
}
