package Users
import (
	"fmt"
	"context"
	"goChore/models"
	"go.mongodb.org/mongo-driver/mongo"
)
func Insert(coll *mongo.Collection) {
	// Insert a single document
	p1 := models.User{
		Id:    1,
		Name:  "John",
		Email: "JOHN@GMAIL.com",
	}
	p2 := models.User{
		Id:    2,
		Name:  "Aman",
		Email: "a@gmail.com",
	}
	p3 := models.User{
		Id:    3,
		Name:  "Jane",
		Email: "Jane@gmail.com",
	}
	// Insert multiple documents
	multiUsers := []interface{}{p1, p2, p3}
	result, err := coll.InsertMany(context.TODO(), multiUsers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Number of documents inserted: %d\n", len(result.InsertedIDs))

}