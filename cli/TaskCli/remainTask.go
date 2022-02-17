package TaskCli

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/api/Tasks/handlerTasks"
	"os"
)

func RemainTask(user *mongo.Collection, task *mongo.Collection) {
	command := os.Args[1]
	username := os.Args[2]
	if command == "remain" {
		handlerTasks.RemainTaskHandler(user, task, username)
		return
	}

}
