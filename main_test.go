package main

import (
	"fmt"
	"testing"

	"github.com/ipfs/testground/cmd"
	"github.com/urfave/cli"
)

func TestMain(t *testing.T) {
	fmt.Println("Jim1")
	app := cli.NewApp()
	app.Name = "testground"
	app.Commands = cmd.Commands
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "v",
			Usage: "verbose output (equivalent to INFO log level)",
		},
		cli.BoolFlag{
			Name:  "vv",
			Usage: "super verbose output (equivalent to DEBUG log level)",
		},
	}
	// Disable the built-in -v flag (version), to avoid collisions with the
	// verbosity flags.
	// TODO implement a `testground version` command instead.
	app.HideVersion = true

	args := []string{
		"./testground",
		"-vv",
		"run",
		"--builder",
		"docker:go",
		"--runner",
		"local:docker",
		"--instances=2",
		"dht/lookup-peers",
	}
	err := app.Run(args)
	if err != nil {
		fmt.Println(err)
	}
}
