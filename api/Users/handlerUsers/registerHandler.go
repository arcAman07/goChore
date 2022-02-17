package handlerUsers

import (
	"context"
	"fmt"
	"goChore/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterHandler(coll *mongo.Collection,username string, password string, status int)  {
	Username := username
	Password := password
	LoggedIn := status
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)

	// Check if the username already exists, if it does we prompt them to change it
	// 0 -> username already exists, 1 -> username does not exist

	// Also check whether the user already exists in the database
	// If yes then prompt then to login rather than register
	// 1 -> user already exists, 0-> user does not exist


	checkUsername := 1
	registeredUser := 1
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in hashing password")
		return
	}
	user := models.User{
		Username: Username,
		Password: Password,
		HashedPassword: (string)(hashedPassword),
		LoggedIn:LoggedIn,
	}
	filter := bson.M{"Username":Username}
	_ ,err = coll.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in finding user")
		return
	}
	if len(password) < 8 {
		fmt.Println("Password must be atleast 8 characters long")
	}
	_ , err = coll.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Registration failed")
		return
	}
	fmt.Println("Registration successful")
	
}