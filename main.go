package main

import (
	"os"

	"github.com/sc-cli/commands"
)

func main() {
	os.Exit(commands.NewApp(os.Args[1:]))
}
