package main

import (
	"github.com/urfave/cli"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "koesie10"
	app.Usage = "Run `status` to get the list of cars"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
