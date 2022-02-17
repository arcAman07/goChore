package UserCli

import (
	"fmt"
	"os"
)

func LogoutUser() (string, string) {
	command := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	if command == "logout" {
		fmt.Println("Logging out user")
		return username, password
	}
	return "", ""
}
