package models
import (
	"time"
	"fmt"
)
var currentTime time.Time = time.Now()
var date string = currentTime.Format("2006-01-02")
fmt.Println(date)
type Task struct {
	ID        string `json:"id,omitempty" bson:"id,omitempty"`
	TaskName string `json:"taskName,omitempty" bson:"taskName,omitempty"`
	Date string `json:"date,omitempty" bson:"date,omitempty"`
	Status bool `json:"status,omitempty" bson:"status,omitempty"`
}