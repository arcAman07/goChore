package handlerTasks

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
	// "image/color"
)
type Gray16 struct {
	Y uint16
}
type Alpha16 struct {
	A uint16
}

var (
	Black       = Gray16{0}
	White       = Gray16{0xffff}
	Transparent = Alpha16{0}
	Opaque      = Alpha16{0xffff}
)

func ListTaskHandler(user *mongo.Collection, task *mongo.Collection, username string) {
	var Username string = username
	Userfilter := bson.M{"Username": Username}
	Taskfilter := bson.M{}
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
		if user.LoggedIn == 1 {
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
				fmt.Printf("%+v\n", task.TaskName)
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
		} else {
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
