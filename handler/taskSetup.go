package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/cli/TaskCli"
	"os"
)

func TaskSetup(user *mongo.Collection, task *mongo.Collection) {
	TaskCli.AddTask(user, task)
	TaskCli.DeleteTask(user, task)
	TaskCli.FinishTask(user, task)
	TaskCli.RemainTask(user, task)
	TaskCli.RefreshTask(user, task)
	TaskCli.UpdateTask(user, task)
	TaskCli.ListTask(user, task)
	if os.Args[1] == "add" {
		UserCli.RegisterUser(user)
	}
	if os.Args[1] == "delete" {
		UserCli.LoginUser(user)
	}
	if os.Args[1] == "finish" {
		UserCli.LogoutUser(user)
	}
	if os.Args[1] == "remain" {
		UserCli.RegisterUser(user)
	}
	if os.Args[1] == "update" {
		UserCli.LoginUser(user)
	}
	if os.Args[1] == "list" {
		UserCli.LogoutUser(user)
	}
}
