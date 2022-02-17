package TaskCli

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/api/Tasks/handlerTasks"
	"os"
)

func ListTask(user *mongo.Collection, task *mongo.Collection) {
	command := os.Args[1]
	username := os.Args[2]
	if command == "list" {
		handlerTasks.ListTaskHandler(user, task, username)
		return
	}
	return
}
