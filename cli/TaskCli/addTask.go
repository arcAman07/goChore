package TaskCli

import (
	"fmt"
	"goChore/cli/UserCli"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddTask(task *mongo.Collection) {
	command := os.Args[1]
	taskName := os.Args[2]
	if command == "add" {
		fmt.Println("Adding task ...")
		return
	}
	return 
}