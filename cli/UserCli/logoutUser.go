package UserCli

import (
	"fmt"
	"os"
)

func LogoutUser() (int) {
	command := os.Args[1]
	if command == "logout" {
		loggedIn = 0
		fmt.Println("Logging out user")
		return loggedIn
	}
	return 2
}
