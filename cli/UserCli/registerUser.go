package UserCli

import (
	"fmt"
	"os"
) 
// 0 means used is logged out
// 1 means the user is logged in
var loggedIn int = 0

func RegisterUser() (string, string, int) {
	command := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	if command == "register" {
		loggedIn = 1
		fmt.Println("Registering user")
		return username, password,loggedIn
	}
	return "", "",2
}
