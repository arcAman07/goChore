package UserCli

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/api/Users/handlerUsers"
	"os"
)

// 0 means used is logged out
// 1 means the user is logged in

func LoginUser(user *mongo.Collection) {
	command := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	if command == "login" {
		fmt.Println("Logging in user ...")
		handlerUsers.LoginHandler(user, username, password)
	}
}
