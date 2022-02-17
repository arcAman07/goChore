package main

// Designing cli apps
import (
	"fmt"
	"os"
	// "goChore/api/Users"
	// "goChore/api/Tasks"
	// "goChore/config"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/cli/UserCli"
)

var UserCollection *mongo.Collection
var TaskCollection *mongo.Collection
// func init() {
// 	UserCollection,TaskCollection = config.Configure()
// 	Users.Insert(UserCollection)
// 	Tasks.Insert(TaskCollection)
// }

func main() {
	username := os.Args[2]
	password := os.Args[3]
	fmt.Println(username, password)
	a,b := UserCli.RegisterUser()
	a,b = UserCli.LoginUser()
	a,b = UserCli.LogoutUser()
	fmt.Println(a,b)
}
