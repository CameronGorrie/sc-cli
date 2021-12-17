package main

import (
	"os"

	cmd "github.com/CameronGorrie/scc/internal"
)

func main() {
	os.Exit(cmd.NewApp(os.Args[1:]))
}
