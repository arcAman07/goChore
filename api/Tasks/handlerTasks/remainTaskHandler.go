package handlerTasks
import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
	// "image/color"
)
func RemainTaskHandler(user *mongo.Collection,task *mongo.Collection, username string) {
	var Username string = username
}