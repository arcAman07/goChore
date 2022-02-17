package TaskCli

import (
	"fmt"
	"goChore/cli/UserCli"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateTask(task *mongo.Collection) (string, string) {
	command := os.Args[1]
	currentTask := os.Args[2]
	newTask := os.Args[3]
	status := UserCli.LoggedIn
	if command == "update" && status == 1 {
		fmt.Println("Updating the task...")
		return currentTask, newTask
	}
	return "", ""
}