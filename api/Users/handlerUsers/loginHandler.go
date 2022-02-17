package handlerUsers
import (
	"context"
	"fmt"
	"goChore/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)
func LoginHandler(coll *mongo.Collection, username string, password string, status int) {
	Username := username
	EnteredPassword := password
	_ = status
	filter := bson.M{"Username":Username}
	update := bson.M{"$set": bson.M{"LoggedIn":1}}
	cursor,err := coll.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
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
			user.LoggedIn = 1
			_ , err := coll.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error in logging in the user")
				fmt.Println("Please try again")
				return
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