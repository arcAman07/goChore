package UserCli

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/api/Users/handlerUsers"
	"os"
)

// 0 means used is logged out
// 1 means the user is logged in

var LoggedIn int

func RegisterUser(coll *mongo.Collection) {
	command := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	if command == "register" {
		LoggedIn = 0
		fmt.Println("Registering user ...")
		handlerUsers.RegisterHandler(coll, username, password,LoggedIn)
		return
	}
}
