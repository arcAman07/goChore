package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/cli/TaskCli"
)
func TaskSetup(task *mongo.Collection) {
	TaskCli.AddTask(task)
	TaskCli.DeleteTask(task)
	TaskCli.FinishTask(task)
	TaskCli.RemainTask(task)
	TaskCli.RefreshTask(task)
	TaskCli.UpdateTask(task)
	TaskCli.ListTask(task)
}