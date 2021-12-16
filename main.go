package main

import (
	"os"

	"github.com/CameronGorrie/scc/cmd"
)

func main() {
	os.Exit(cmd.NewApp(os.Args[1:]))
}
