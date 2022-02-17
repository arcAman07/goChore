package handlerTasks

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
)

func AddTaskHandler(coll *mongo.Collection, username string, task string) {
	var Username string = username
	var Task string = task
	filter := bson.M{"Username": Username}
	cursor, err := coll.Find(context.TODO(), filter)
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
			fmt.Println("Adding task ...")
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
