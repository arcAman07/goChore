package Tasks

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"goChore/models"
	"time"
)

var currentTime time.Time = time.Now()
var currentDay string = currentTime.Format("01-02-2006")

func Insert(coll *mongo.Collection) {
	// Insert a single document
	p1 := models.Task{
		TaskName: "study physics",
		Date:     currentDay,
		Status:   "Not Done",
		UserName: "Deepak",
	}
	p2 := models.Task{
		TaskName: "code in golang",
		Date:     currentDay,
		Status:   "Not Done",
		UserName: "Aman",
	}
	p3 := models.Task{
		TaskName: "Eating Food",
		Date:     currentDay,
		Status:   "Done",
		UserName: "Roma",
	}
	// Insert multiple documents
	multiTasks := []interface{}{p1, p2, p3}
	result, err := coll.InsertMany(context.TODO(), multiTasks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Number of documents inserted: %d\n", len(result.InsertedIDs))

}
