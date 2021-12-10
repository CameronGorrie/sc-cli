package main

import (
	"os"

	"github.com/sc-cli/cmd"
)

func main() {
	os.Exit(cmd.NewApp(os.Args[1:]))
}
