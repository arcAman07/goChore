package TaskCli

import (
	"fmt"
	"goChore/cli/UserCli"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteTask(user *mongo.Collection,task *mongo.Collection) {
	command := os.Args[1]
	userName := os.Args[2]
	taskName := os.Args[3]
	if command == "delete" {
		handlerTasks.AddTaskHandler(user,task,userName, taskName)
		return
	}
	return 
}