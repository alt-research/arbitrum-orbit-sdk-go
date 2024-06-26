package commands

import (
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func init() {
	app.Name = "arbctl"
	app.Usage = "Cli tool to interact with the Arbitrum chain"
	app.Commands = []cli.Command{}
}

func Run(args ...string) error {
	if len(args) == 0 {
		args = os.Args
	}
	return app.Run(args)
}
