package TaskCli

import (
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/api/Tasks/handlerTasks"
)

func RefreshTask(user *mongo.Collection,task *mongo.Collection) {
	command := os.Args[1]
	userName := os.Args[2]
	taskName := os.Args[3]
	if command == "refresh" {
		handlerTasks.RefreshTaskHandler(user,task,userName, taskName)
		return
	}
	return 
}
