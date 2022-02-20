package main

// Designing cli apps

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/config"
	"goChore/handler"
)

var UserCollection *mongo.Collection
var TaskCollection *mongo.Collection

func init() {

}

func main() {

	UserCollection, TaskCollection = config.Configure()
	handler.UserSetup(UserCollection)
	handler.TaskSetup(UserCollection, TaskCollection)
	handler.Helper()

}


// Better auth techniques to be added ( tokens ), CI-CD pipelines

