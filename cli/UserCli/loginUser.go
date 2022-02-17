package UserCli

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func LoginUser(coll *mongo.Collection) (string, string, int) {
	command := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	if command == "login" && LoggedIn == 0 {
		LoggedIn = 1
		fmt.Println("Logging in user")
		return username, password,LoggedIn
	}
	if command == "login" && LoggedIn == 1 {
		fmt.Println("User is already logged in")
		return "", "",2
	}
	return "", "",2
}