package initializeUsers
import (
	"fmt"
	"context"
	"goChore/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// LoggedIn in the schema tells us whether the user is logged in or not

// 0 -> loggedOut, 1 -> loggedIn

func Insert(coll *mongo.Collection) {
	// Insert a single document
	hash, err := bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	p1 := models.User{
		Username: "arcAman07",
		Password: "12345",
		HashedPassword:(string)(hash),
		LoggedIn:0,
	}
	hash, err = bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	p2 := models.User{
		Username: "john05",
		Password: "123456",
		HashedPassword: (string)(hash),
		LoggedIn:1,
	}
	hash, err = bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	p3 := models.User{
		Username: "janeb4",
		Password: "1234567",
		HashedPassword: (string)(hash),
		LoggedIn:1,
	
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