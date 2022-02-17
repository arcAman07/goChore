package TaskCli
import (
	"fmt"
	"os"
	"goChore/cli/UserCli"
)

func ListTask() (string, string) {
	command := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	if command == "login" {
		fmt.Println("Logging in user")
		return username, password
	}
	return "", ""
}