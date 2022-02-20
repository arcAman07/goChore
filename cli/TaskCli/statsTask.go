package TaskCli

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/api/Tasks/handlerTasks"
	"os"
)

func StatsTask(user *mongo.Collection, task *mongo.Collection) {
	command := os.Args[1]
	userName := os.Args[2]
	if command == "stats" {
		handlerTasks.StatsTaskHandler(user, task, userName)
		return
	}
}
