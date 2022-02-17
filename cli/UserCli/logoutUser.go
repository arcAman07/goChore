package UserCli

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/api/Users/handlerUsers"
	"os"
)

// 0 means used is logged out
// 1 means the user is logged in

func LogoutUser(coll *mongo.Collection) {
	command := os.Args[1]
	username := os.Args[2]
	if command == "logout" {
		fmt.Println("Logging out user ...")
		handlerUsers.LogoutHandler(coll, username)
		return
	}
}
