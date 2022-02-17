package TaskCli

import (
	"fmt"
	"goChore/cli/UserCli"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func RefreshTask(task *mongo.Collection) {
	command := os.Args[1]
	status := UserCli.LoggedIn
	if command == "refresh" && status == 1 {
		fmt.Println("Refreshing all tasks...")
	}
}