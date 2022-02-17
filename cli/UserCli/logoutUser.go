package UserCli

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func LogoutUser(coll *mongo.Collection) (int) {
	command := os.Args[1]
	if command == "logout" {
		LoggedIn = 0
		fmt.Println("Logging out user")
		return LoggedIn
	}
	return 2
}
