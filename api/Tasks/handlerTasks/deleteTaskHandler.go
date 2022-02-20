package handlerTasks

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
)

func DeleteTaskHandler(user *mongo.Collection, task *mongo.Collection, username string, taskName string) {
	var Username string = username
	var Task string = taskName
	Userfilter := bson.M{"Username": Username}
	cursor, err := user.Find(context.TODO(), Userfilter)
	Taskfilter := bson.M{"TaskName": Task}
	if err != nil {
		fmt.Println("Username does not exist")
		fmt.Println("Please register")
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
			fmt.Println("Deleting task ...")
			_, err = task.DeleteOne(context.TODO(), Taskfilter)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error in deleting task")
				fmt.Println("Please check task name")
				return
			}
			fmt.Println("Task deleted successfully")
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
