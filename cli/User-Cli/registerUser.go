package UserCli

import (
	"fmt"
	"os"
)

func RegisterUser() (string, string) {
	command := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	if command == "register" {
		fmt.Println("Registering user")
		return username, password
	}
	return "", ""
}
