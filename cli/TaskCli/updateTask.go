package TaskCli

import (
	
	"goChore/api/Tasks/handlerTasks"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateTask(user *mongo.Collection,task *mongo.Collection) {
	command := os.Args[1]
	userName := os.Args[2]
	currentTask := os.Args[3]
	newTask := os.Args[4]
	if command == "update"{
		handlerTasks.UpdateTaskHandler(user,task,userName,currentTask,newTask)
		return
	}
}