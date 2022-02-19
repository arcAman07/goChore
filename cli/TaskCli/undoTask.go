package TaskCli

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/api/Tasks/handlerTasks"
	"os"
)

func UndoTask(user *mongo.Collection, task *mongo.Collection) {
	command := os.Args[1]
	userName := os.Args[2]
	currentTask := os.Args[3]
	if command == "undo" {
		handlerTasks.UndoTaskHandler(user, task, userName, currentTask)
		return
	}
}