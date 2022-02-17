package TaskCli
import (
	"fmt"
	"os"
	"goChore/cli/UserCli"
)

func RefreshTask() {
	command := os.Args[1]
	status := UserCli.LoggedIn
	if command == "refresh" && status == 1 {
		fmt.Println("Refreshing all tasks...")
	}
}