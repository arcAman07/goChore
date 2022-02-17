package TaskCli
import (
	"fmt"
	"os"
	"goChore/cli/UserCli"
)

func FinishTask() (string) {
	command := os.Args[1]
	taskName := os.Args[2]
	status := UserCli.LoggedIn
	if command == "finish" && status == 1 {
		fmt.Println("Finishing task")
		return taskName
	}
	fmt.Println("Not logged in")
	return ""
}