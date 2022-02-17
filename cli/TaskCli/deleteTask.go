package TaskCli
import (
	"fmt"
	"os"
	"goChore/cli/UserCli"
)

func DeleteTask() (string) {
	command := os.Args[1]
	taskName := os.Args[2]
	status := UserCli.LoggedIn
	if command == "delete" && status == 1 {
		fmt.Println("Deleting task")
		return taskName
	}
	fmt.Println("Not logged in")
	return ""
}