package main

import (
	"fmt"
	"os"

	cmd "github.com/CameronGorrie/scc/internal"
)

func main() {
	if len(os.Args) <= 1 {
		helpMsg := `NAME:
	scc - Supercollider CLI
 
USAGE:
	scc [global options] command [command options] [arguments...]
 
COMMANDS:
	free	Frees nodes in a running instance of scsynth
	send	Creates and sends synthdefs to a running instance of scsynth
 
GLOBAL OPTIONS:
	-h	Show help for a command
	-u	The remote address for scsynth
`

		fmt.Println(helpMsg)
		os.Exit(0)
	}

	os.Exit(cmd.NewApp(os.Args[1:]))
}
