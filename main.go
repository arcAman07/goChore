package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/arcAman07/Go-Tasks/actions"
)

// Designing cli apps
var app = 

func main() {
	app.Name = "Go-Tasks"

	app.Commands = []cli.Command{
		{
			Name: "noop",
		},
		{
			Name:     "add",
			Category: "template",
		},
		{
			Name:     "remove",
			Category: "template",
		},
	}

	app.Run(os.Args)
}
