package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/baulk/baulk-go/command"
)

type options struct {
	CMD string
}

func (opt *options) Invoke(val int, oa, raw string) error {
	switch val {
	case int('v'):
		fmt.Fprintf(os.Stderr, "1.0\n")
		os.Exit(0)
	case 'h':
		fmt.Fprintf(os.Stderr, "usage: option\n")
		os.Exit(0)
	case 'c':
		opt.CMD = oa
	default:
		return errors.New("unsupport val")
	}
	return nil
}

func main() {
	var ae command.ArgvEngine
	ae.Add("version", command.NOARG, 'v')
	ae.Add("help", command.NOARG, 'h')
	ae.Add("cmd", command.REQUIRED, 'c')
	var opt options
	if err := ae.Execute(os.Args, &opt); err != nil {
		fmt.Fprintf(os.Stderr, "ParseArgv: \x1b[31m%s\x1b[0m\n", err.Error())
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "cmd: %s\nUnresolved: %v\n", opt.CMD, ae.Unresolved())
}
