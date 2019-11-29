package command

import (
	"fmt"
	"os"
)

// baulk -E command mode

func executeUsage() {

}

// Exec todo
func Exec(args []string) int {
	//
	switch args[0] {
	case "hash":
		return HashCalculate(args)
	case "wget":
		return WebGet(args)
	case "tar":
		return BaulkTar(args)
	case "-h", "--help":
		executeUsage()
		os.Exit(0)
	}
	fmt.Fprintf(os.Stderr, "unregister command '%s'\n", args[0])
	return 0
}
