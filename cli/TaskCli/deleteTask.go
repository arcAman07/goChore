package TaskCli

import (
	"fmt"
	"goChore/cli/UserCli"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteTask(task *mongo.Collection) (string) {
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