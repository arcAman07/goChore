package main

// Designing cli apps
import (
	"fmt"
	"os"
	// "goChore/api/Users"
	// "goChore/api/Tasks"
	// "goChore/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection
var TaskCollection *mongo.Collection
// func init() {
// 	UserCollection,TaskCollection = config.Configure()
// 	Users.Insert(UserCollection)
// 	Tasks.Insert(TaskCollection)
// }

func main() {
	username := os.Args[0]
	password := os.Args[1]
	fmt.Println(username, password)
}
