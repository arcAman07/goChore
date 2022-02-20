package handlerUsers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
)

func LoginHandler(User *mongo.Collection, username string, password string) {
	Username := username
	EnteredPassword := password
	filter := bson.M{"Username": Username}
	update := bson.M{"$set": bson.M{"LoggedIn": "1"}}
	anotherFilter := bson.D{{"Username", username}}
	cursor, err := User.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = User.FindOne(context.TODO(), anotherFilter).Decode(&User)
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
		Password := user.Password
		if Password == EnteredPassword {
			fmt.Println("Logged in successfully")
			if user.LoggedIn == "1" {
				fmt.Println("You are already logged in")
				return
			}
			if user.LoggedIn == "0" {
				user.LoggedIn = "1"
				_, err := User.UpdateOne(context.TODO(), filter, update)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Error in logging in the user")
					fmt.Println("Please try again")
					return
				}
			}
		}
		if Password != EnteredPassword {
			fmt.Println("Incorrect password")
			fmt.Println("Please try again")
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
