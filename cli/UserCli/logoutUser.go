package UserCli

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

// 0 means used is logged out
// 1 means the user is logged in

func LogoutUser(coll *mongo.Collection) (string,int) {
	command := os.Args[1]
	username := os.Args[2]
	if command == "logout" && LoggedIn == 1 {
		LoggedIn = 0
		fmt.Println("Logging out user ...")
		return username,LoggedIn
	}
	if command == "logout" && LoggedIn == 0 {
		fmt.Println("User is already logged out")
		fmt.Println("Please login again")
		return "",2
	}
	return "",2
}
