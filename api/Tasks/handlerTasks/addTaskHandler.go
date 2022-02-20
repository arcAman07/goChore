package handlerTasks

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
	"time"
)

var currentTime time.Time = time.Now()
var currentDay string = currentTime.Format("01-02-2006")

func AddTaskHandler(user *mongo.Collection, task *mongo.Collection, userName string, taskName string) {
	var username string = userName
	var Task string = taskName
	filter := bson.M{"Username":username}
	cursor, err := user.Find(context.TODO(), filter)
	anotherFilter := bson.D{{"Username", username}}
	if err != nil {
		fmt.Println("Username does not exist")
		fmt.Println("Please register")
		return
	}
	err = user.FindOne(context.TODO(), anotherFilter).Decode(&user)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in decoding user")
		return
	}
	// Iterate through the cursor
	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error in decoding user")
			return
		}
		if user.LoggedIn == "1" {
			fmt.Println("Adding task ...")
			enterTask := models.Task{
				TaskName: Task,
				Date:     currentDay,
				Status:   "Not Done",
				UserName: username,
			}
			_, err = task.InsertOne(context.TODO(), enterTask)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error in inserting task")
				return
			}
			fmt.Println("Task added successfully")
		} 
		if user.LoggedIn == 0 {
			fmt.Println("Please login first")
			return
		}
		if err := cursor.Err(); err != nil {
			fmt.Println(err)
			fmt.Println("Error in cursor")
			return
		}
		// Close the cursor once finished
		if err := cursor.Close(context.TODO()); err != nil {
			fmt.Println(err)
			fmt.Println("Error in closing cursor")
			return
		}

	}
}
