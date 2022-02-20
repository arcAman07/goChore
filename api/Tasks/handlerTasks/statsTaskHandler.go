package handlerTasks

import (
	"context"
	"fmt"
	"goChore/models"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func StatsTaskHandler(user *mongo.Collection, task *mongo.Collection, username string) {
	var Username string = username
	Userfilter := bson.M{"Username": Username}
	cursor, err := user.Find(context.TODO(), Userfilter)
	var doneTasks int = 0
	var pendingTasks int = 0
	Taskfilter := bson.M{}
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
			fmt.Println("Updating task status ...")
			cursor, err = task.Find(context.TODO(), Taskfilter)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error in updating task status")
				fmt.Println("Please check task name")
				return
			}
			for cursor.Next(context.TODO()) {
				var task models.Task
				err := cursor.Decode(&task)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Error in decoding task")
					return
				}
				if task.Status == "Done" {
					doneTasks++
				}
				if task.Status == "Not Done" {
					pendingTasks++
				}

			}
			if err := cursor.Err(); err != nil {
				fmt.Println(err)
				return
			}
			// Close the cursor once finished
			if err := cursor.Close(context.TODO()); err != nil {
				fmt.Println(err)
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
		color.Green("Done tasks for the day is %d", doneTasks)
		color.Red("Pending tasks for the day is %d", pendingTasks)
		var total int = doneTasks + pendingTasks
		var ratio float64 = (float64(total)/float64(doneTasks))*100
		color.Blue("Percentage of done tasks to all tasks is %f", ratio)
		if ratio > 75 {
			color.Green("You are doing very well. Keep it up")

		} 
		if ratio < 75 && ratio > 50 {
			color.Yellow("You are doing well. Keep it up")
		}
		if ratio < 50 {
			color.Red("You are doing poorly. You need to work hard.")

		}

		// Add other stats also if needed be( comments based on ratio )

	}
}
