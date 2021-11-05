package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "register",
			Aliases: []string{"r"},
			Usage:   "Register a new user",
			Action: func(c *cli.Context) error {
				fmt.Println("Registering a new user")
				return nil
			},
		},
	}
}
