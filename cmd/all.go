package cmd

import (
	"github.com/ipfs/testground/pkg/engine"
	"github.com/urfave/cli"
)

// _engine is a singleton engine to be shared by all CLI commands.
var _engine = engine.NewDefaultEngine()

// Commands collects all subcommands of the testground CLI.
var Commands = []cli.Command{
	RunCommand,
	ListCommand,
	BuildCommand,
}
