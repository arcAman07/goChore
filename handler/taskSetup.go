package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/cli/TaskCli"
)
func TaskSetup(task *mongo.Collection) {
	TaskCli.AddTask(user,task)
	TaskCli.DeleteTask(user,task)
	TaskCli.FinishTask(user,task)
	TaskCli.RemainTask(user,task)
	TaskCli.RefreshTask(user,task)
	TaskCli.UpdateTask(task)
	TaskCli.ListTask(task)
}