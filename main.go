package main

import (
	"os"

	"github.com/sc-cli/commands"
)

func main() {
	os.Exit(commands.CLI(os.Args[1:]))
}
