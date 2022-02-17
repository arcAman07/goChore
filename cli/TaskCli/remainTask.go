package TaskCli
import (
	"fmt"
	"os"
	"goChore/cli/UserCli"
)

func RemainTask() {
	command := os.Args[1]
	status := UserCli.LoggedIn
	if command == "login" && status == 1 {
		fmt.Println("Logging in user")
	}
}