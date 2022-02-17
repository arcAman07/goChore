package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/cli/TaskCli"
	"os"
)

func TaskSetup(user *mongo.Collection, task *mongo.Collection) {
	if os.Args[1] == "add" {
		TaskCli.AddTask(user, task)
	}
	if os.Args[1] == "delete" {
		TaskCli.DeleteTask(user, task)
	}
	if os.Args[1] == "finish" {
		TaskCli.FinishTask(user, task)
	}
	if os.Args[1] == "remain" {
		TaskCli.RemainTask(user, task)
	}
	if os.Args[1] == "update" {
		TaskCli.UpdateTask(user, task)
	}
	if os.Args[1] == "list" {
		TaskCli.ListTask(user, task)
	}
	if os.Args[1] == "refresh" {
		TaskCli.RefreshTask(user, task)
	}
}
