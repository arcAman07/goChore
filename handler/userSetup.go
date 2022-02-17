package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/cli/UserCli"
)
func UserSetup(user *mongo.Collection) {
	UserCli.RegisterUser(user)
	UserCli.LoginUser(user)
	UserCli.LogoutUser(user)
}