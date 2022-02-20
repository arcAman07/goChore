package handlerTasks

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
)

// Red Color denotes the task is "Not Done"

// Green Color denotes the task is "Done"

func ListTaskHandler(user *mongo.Collection, task *mongo.Collection, username string) {
	var Username string = username
	Userfilter := bson.M{"Username": Username}
	Taskfilter := bson.M{}
	var i int = 1
	cursor, err := user.Find(context.TODO(), Userfilter)
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
			fmt.Println("Listing tasks ...")
			cursor, err = task.Find(context.TODO(), Taskfilter)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error in finding tasks")
				return
			}
			// Iterate through the cursor
			for cursor.Next(context.TODO()) {
				var task models.Task
				err := cursor.Decode(&task)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Error in decoding task")
					return
				}
				if task.Status == "Not Done" {
					color.Red("%d\t%+v\n", i, task.TaskName)
					fmt.Println()
					i++
				}
				if task.Status == "Done" {
					color.Green("%d\t%+v\n", i, task.TaskName)
					fmt.Println()
					i++
				}
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
