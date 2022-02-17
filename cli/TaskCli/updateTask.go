package TaskCli

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/api/Tasks/handlerTasks"
	"os"
)

func UpdateTask(user *mongo.Collection, task *mongo.Collection) {
	command := os.Args[1]
	userName := os.Args[2]
	currentTask := os.Args[3]
	newTask := os.Args[4]
	if command == "update" {
		handlerTasks.UpdateTaskHandler(user, task, userName, currentTask, newTask)
		return
	}
}
