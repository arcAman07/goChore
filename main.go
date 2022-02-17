package main

// Designing cli apps

import (
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/config"
	"goChore/handler"
)

var UserCollection *mongo.Collection
var TaskCollection *mongo.Collection

func main() {
	UserCollection, TaskCollection = config.Configure()
	handler.UserSetup(UserCollection)
	handler.TaskSetup(UserCollection,TaskCollection)
}

// func init will be used to populate some raw dummy data in the database

// Have to be implemented

// Write test code for testing the functions, and implement logger

// Better auth techniques to be added ( tokens ), CI-CD pipelines

// better errer handling and use of go routines

