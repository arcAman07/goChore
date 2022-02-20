package handlerUsers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(coll *mongo.Collection, username string, password string) {
	Username := username
	Password := password
	var currentLog int = 0
	// Check if the username already exists, if it does we prompt them to change it
	// 0 -> username does not exists, 1 -> username already exists

	// Also check whether the user already exists in the database
	// If yes then prompt then to login rather than register
	// 0-> user does not exist, 1 -> user already exists,

	var checkUsername int = 0
	var registeredUser int = 0

	if len(password) < 8 {
		fmt.Println("Password must be atleast 8 characters long")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in hashing password")
		return
	}

	// Check if the user already exists in the database

	// Check if the username already exists
	filter := bson.M{"Username": Username}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		checkUsername = 0
		registeredUser = 0
	} else {
		for cursor.Next(context.TODO()) {
			var user models.User
			err := cursor.Decode(&user)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error in decoding user")
				return
			}
			CorrectPassword := user.Password
			err = bcrypt.CompareHashAndPassword([]byte(CorrectPassword), []byte(Password))
			if err != nil {
				// Check if the username already exists
				fmt.Println("Username already in database")
				checkUsername = 1
				registeredUser = 0
				fmt.Println("Please choose a different username")
				return
			} else {
				// Check if the user already exists in the database
				checkUsername = 1
				registeredUser = 1
				fmt.Println("User already in database")
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
	// If the username entered is not used and user is not registered, then we can register the user
	if checkUsername == 0 && registeredUser == 0 {
		user := models.User{
			Username:       Username,
			Password:       Password,
			HashedPassword: (string)(hashedPassword),
			LoggedIn:       currentLog,
		}
		_, err = coll.InsertOne(context.TODO(), user)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Registration failed")
			return
		}
		fmt.Println("Registration successful")
	}

}
