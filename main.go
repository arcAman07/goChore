package main

// Designing cli apps
import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/config"
	"goChore/handler"
)

var UserCollection *mongo.Collection
var TaskCollection *mongo.Collection

// func init() {
// 	UserCollection,TaskCollection = config.Configure()
// 	Users.Insert(UserCollection)
// 	Tasks.Insert(TaskCollection)
// }

func main() {
	UserCollection, TaskCollection = config.Configure()
	handler.UserSetup(UserCollection)
	// handler.TaskSetup(TaskCollection)
}
