package main

import (
	"os"

	"github.com/baulk/baulk/command"
)

var (
	version = "0.0.0"
	commit  = "dev"
	date    = "unknown"
)

// -E command
func main() {
	//command.Exec(os.Args)
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	if os.Args[1] == "-E" {
		os.Exit(command.Exec(os.Args))
	}
}
