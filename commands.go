package main

import (
	"fmt"
	"os"

	"github.com/koesie10/assettocorsa/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "status",
		Usage:  "",
		Action: command.CmdStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Value: "localhost",
				Usage: "The host to retrieve the settings for",
			},
			cli.UintFlag{
				Name:  "port",
				Value: 8081,
				Usage: "The port to retrieve the settings for",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
