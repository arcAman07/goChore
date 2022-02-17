package UserCli

import (
	"fmt"
	"os"
)

func LogoutUser() (int) {
	command := os.Args[1]
	if command == "logout" {
		LoggedIn = 0
		fmt.Println("Logging out user")
		return LoggedIn
	}
	return 2
}
