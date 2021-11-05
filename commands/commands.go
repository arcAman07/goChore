package actions

import (
	"fmt"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func createUser() {

	fmt.Println("New User created")
}

func loginUser() {
	fmt.Println("User logged in")
}
