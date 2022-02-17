package TaskCli

import (
	"os"
	"goChore/api/Tasks/handlerTasks"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddTask(user *mongo.Collection,task *mongo.Collection) {
	command := os.Args[1]
	userName := os.Args[2]
	taskName := os.Args[3]
	if command == "add" {
		handlerTasks.AddTaskHandler(user,task,userName, taskName)
		return
	}
}