package Users
import (
	"fmt"
	"context"
	"goChore/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)
func Insert(coll *mongo.Collection) {
	// Insert a single document
	p1 := models.User{
		Username: "arcAman07",
		Password: "12345",
		HashedPassword:bcrypt.GenerateFromPassword([]byte("12345"), bcrypt.DefaultCost)
	}
	p2 := models.User{
		Username: "john05",
		Password: "123456",
		HashedPassword: bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	}
	p3 := models.User{
		Username: "janeb4",
		Password: "1234567",
		HashedPassword: bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
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