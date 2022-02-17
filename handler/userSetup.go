package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/cli/UserCli"
	"os"
)

func UserSetup(user *mongo.Collection) {
	if os.Args[1] == "register" {
		UserCli.RegisterUser(user)
	}
	if os.Args[1] == "login" {
		UserCli.LoginUser(user)
	}
	if os.Args[1] == "logout" {
		UserCli.LogoutUser(user)
	}
}
