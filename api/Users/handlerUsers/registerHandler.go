package handlerUsers

import (
	"context"
	"fmt"
	"goChore/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(coll *mongo.Collection,username string, password string, status int)  {
	Username := username
	Password := password
	LoggedIn := status
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
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
	_ , err = coll.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Registration failed")
		return
	}
	fmt.Println("Registration successful")
	
}