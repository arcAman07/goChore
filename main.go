package main

// Designing cli apps
import (
	"goChore/actions"
	"fmt"
	"log"
	"os"
	"goChore/api/Users"
	"goChore/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var app = actions.App
var UserCollection *mongo.Collection
var TaskCollection *mongo.Collection
func init() {
	UserCollection,TaskCollection = config.Configure()
	Users.Insert(UserCollection)
}

func main() {
	fmt.Println("Hello World")
	actions.Actions()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
