package initializeUsers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
)

func Find(coll *mongo.Collection) {
	// Find all documents
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Iterate through the cursor
	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", user)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println(err)
		return
	}
	// Close the cursor once finished
	if err := cursor.Close(context.TODO()); err != nil {
		fmt.Println(err)
		return
	}
}
