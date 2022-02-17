package handlerUsers

import (
	"context"
	"fmt"
	"goChore/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func LogoutHandler(coll *mongo.Collection, username string,status int) {
	Username := username
	LoggedIn := status
	filter := bson.M{"Username":Username}
	cursor,err := coll.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Please enter your valid username")
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
		user.LoggedIn = LoggedIn
		fmt.Println("Logged out successfully")
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