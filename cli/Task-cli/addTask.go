package TaskCli
import (
	"fmt"
	"os"
)

func AddTask() (string) {
	command := os.Args[1]
	username := os.Args[2]
	if command == "login" {
		fmt.Println("Logging in user")
		return username, password
	}
	return "", ""
}