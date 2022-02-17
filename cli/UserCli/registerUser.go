package UserCli

import (
	"fmt"
	"os"
	"goChore/api/Users/handlerUsers"
	"go.mongodb.org/mongo-driver/mongo"
)

// 0 means used is logged out
// 1 means the user is logged in
var LoggedIn int = 0

func RegisterUser(coll *mongo.Collection) {
	command := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	if command == "register" && LoggedIn == 0{
		LoggedIn = 1
		fmt.Println("Registering user ...")
		handlerUsers.RegisterHandler(coll, username, password, LoggedIn)
	}
	if command == "register" && LoggedIn == 1 {
		fmt.Println("User is already logged in")
	}
}
