package initializeUsers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(coll *mongo.Collection) {
	// Delete a single document
	filter := bson.M{"Username": "john05"}
	res, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Number of documents deleted: %d\n", res.DeletedCount)
}
