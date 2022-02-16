package Users
import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)
func Update(coll *mongo.Collection){
	// Update a single document
	filter := bson.M{"Username": "arcAman07"}
	update := bson.M{"$set": bson.M{"Username": "roasterAman07"}}
	resp, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Number of documents updated: %d\n", resp.ModifiedCount)
}