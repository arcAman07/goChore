package UserCli
import (
	"fmt"
	"os"
)

func LoginUser() (string, string, int) {
	command := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	if command == "login" {
		loggedIn = 1
		fmt.Println("Logging in user")
		return username, password,loggedIn
	}
	return "", "",2
}