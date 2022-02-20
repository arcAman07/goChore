package handlerTasks

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
)

func UpdateTaskHandler(user *mongo.Collection, task *mongo.Collection, username string, taskName string, updatedTask string) {
	var Username string = username
	var Task string = taskName
	var UpdatedTask string = updatedTask
	Userfilter := bson.M{"Username": Username}
	cursor, err := user.Find(context.TODO(), Userfilter)
	Taskfilter := bson.M{"TaskName": Task}
	update := bson.M{"$set": bson.M{"TaskName": UpdatedTask}}
	if err != nil {
		fmt.Println("Username does not exist")
		fmt.Println("Please register")
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
			fmt.Println("Updating task name ...")
			_, err = task.UpdateOne(context.TODO(), Taskfilter, update)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error in updating task name")
				fmt.Println("Please check task name")
			}
			fmt.Println("Task name updated successfully")
		}
		if user.LoggedIn == "0" {
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
