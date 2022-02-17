package TaskCli

import (
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/api/Tasks/handlerTasks"
)

func RefreshTask(user *mongo.Collection,task *mongo.Collection) {
	command := os.Args[1]
	userName := os.Args[2]
	if command == "refresh" {
		handlerTasks.RefreshTaskHandler(user,task,userName)
		return
	}
}
