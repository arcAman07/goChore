package Tasks

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Update(coll *mongo.Collection) {
	// Update a single document
	filter := bson.M{"TaskName": "code in golang"}
	update := bson.M{"$set": bson.M{"status": true}}
	resp, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Number of documents updated: %d\n", resp.ModifiedCount)
}
