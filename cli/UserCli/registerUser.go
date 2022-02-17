package UserCli

import (
	"fmt"
	"os"
) 
// 0 means used is logged out
var loggedIn int = 0

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
