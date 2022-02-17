package TaskCli

import (
	"fmt"
	"os"
	"goChore/api/Tasks/handlerTasks"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddTask(task *mongo.Collection) {
	command := os.Args[1]
	userName := os.Args[2]
	taskName := os.Args[3]
	if command == "add" {
		handlerTasks.AddTaskHandler(task,userName, taskName)
		return
	}
	return 
}