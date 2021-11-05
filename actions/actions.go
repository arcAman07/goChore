package actions

import (
	"fmt"

	"github.com/urfave/cli"
)

var App = cli.NewApp()
var RegisterApp = cli.NewApp()
var LoginApp = cli.NewApp()
var RegisterUsername string
var RegisterPassword string
var LoginUsername string
var LoginPassword string

func Actions() {
	App.Commands = []cli.Command{
		{
			Name:    "register",
			Aliases: []string{"r"},
			Usage:   "Register a new user",
			Action: func(c *cli.Context) (error, string, string) {
				RegisterUsername = c.Args().Get(0)
				RegisterPassword = c.Args().Get(1)
				fmt.Println("Registering a new user")
				fmt.Println(RegisterUsername, '\t', RegisterPassword)
				RegisterApp.Commands = []cli.Command{
					{
						Name:    "add",
						Aliases: []string{"a"},
						Usage:   "Add a new task",
						Action: func(c *cli.Context) error {
							fmt.Println("Adding a new task")
							return nil
						},
					},
					{
						Name:    "get",
						Aliases: []string{"g"},
						Usage:   "Get all tasks",
						Action: func(c *cli.Context) error {
							fmt.Println("Getting all tasks")
							return nil
						},
					},
					{
						Name:    "delete",
						Aliases: []string{"d"},
						Usage:   "Delete a task",
						Action: func(c *cli.Context) error {
							fmt.Println("Deleting a task")
							return nil
						},
					},
				}
				return nil, RegisterUsername, RegisterPassword
			},
		},

		{
			Name:    "login",
			Aliases: []string{"l"},
			Usage:   "Login to your account",
			Action: func(c *cli.Context) (error, string, string) {
				LoginUsername = c.Args().Get(0)
				LoginPassword = c.Args().Get(1)
				fmt.Println("Logging in")
				fmt.Println(LoginUsername, '\t', LoginPassword)
				LoginApp.Commands = []cli.Command{
					{
						Name:    "add",
						Aliases: []string{"a"},
						Usage:   "Add a new task",
						Action: func(c *cli.Context) error {
							fmt.Println("Adding a new task")
							return nil
						},
					},
					{
						Name:    "get",
						Aliases: []string{"g"},
						Usage:   "Get all tasks",
						Action: func(c *cli.Context) error {
							fmt.Println("Getting all tasks")
							return nil
						},
					},
					{
						Name:    "delete",
						Aliases: []string{"d"},
						Usage:   "Delete a task",
						Action: func(c *cli.Context) error {
							fmt.Println("Deleting a task")
							return nil
						},
					},
				}
				return nil, LoginUsername, LoginPassword
			},
		},
		{
			Name:    "logout",
			Aliases: []string{"lg"},
			Usage:   "Logout from your account",
			Action: func(c *cli.Context) error {
				fmt.Println("Logging out")
				return nil
			},
		},
	}
}
