package TaskCli

import (
	"fmt"
	"goChore/cli/UserCli"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddTask(task *mongo.Collection) (string) {
	command := os.Args[1]
	taskName := os.Args[2]
	status := UserCli.LoggedIn
	if command == "add" && status == 1 {
		fmt.Println("Adding task")
		return taskName
	}
	fmt.Println("Not logged in")
	return ""
}