package TaskCli

import (
	"fmt"
	"goChore/cli/UserCli"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func RemainTask(task *mongo.Collection) {
	command := os.Args[1]
	status := UserCli.LoggedIn
	if command == "login" && status == 1 {
		fmt.Println("Logging in user")
	}
}