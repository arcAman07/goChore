package handlerUsers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
)

func LogoutHandler(coll *mongo.Collection, username string) {
	var Username string = username
	filter := bson.M{"Username": Username}
	update := bson.M{"$set": bson.M{"LoggedIn": "0"}}
	cursor, err := coll.Find(context.TODO(), filter)
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
		if user.LoggedIn == "1" {
			fmt.Println("Logged out successfully")
			user.LoggedIn = "0"
			_, err := coll.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error in logging out the user")
				fmt.Println("Please try again")
				return
			}
			return
		}
		if user.LoggedIn == "0" {
			fmt.Println("You are not logged in")
			fmt.Println("Please login")
			return
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
