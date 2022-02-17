package TaskCli
import (
	"fmt"
	"os"
	"goChore/cli/UserCli"
)

func RefreshTask() (string, string) {
	command := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	status := UserCli.LoggedIn
	if command == "login" && status == 1 {
		fmt.Println("Logging in user")
		return username, password
	}
	return "", ""
}