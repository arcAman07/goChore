package handlerTasks

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
)

func RefreshTaskHandler(user *mongo.Collection, task *mongo.Collection, username string) {
	var Username string = username
	Userfilter := bson.M{"Username": Username}
	cursor, err := user.Find(context.TODO(), Userfilter)
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
			fmt.Println("Deleting all tasks ...")
			_, err = task.DeleteMany(context.TODO(), Taskfilter)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error in deleting all tasks")
				fmt.Println("Please try again")
				return
			}
			fmt.Println("All tasks deleted successfully")
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
