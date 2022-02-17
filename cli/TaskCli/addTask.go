package TaskCli
import (
	"fmt"
	"os"
	"goChore/cli/UserCli"
)

func AddTask() (string) {
	command := os.Args[1]
	username := os.Args[2]
	status := UserCli.LoggedIn
	if command == "login" && status == 1 {
		fmt.Println("Logging in user")
		return username
	}
	return ""
}