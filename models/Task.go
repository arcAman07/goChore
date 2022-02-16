package models
import (
	"time"
)
var currentTime time.Time = time.Now()
var Date string = currentTime.Format("01-02-2006")
type Task struct {
	ID        string `json:"id,omitempty" bson:"id,omitempty"`
	TaskName string `json:"taskName,omitempty" bson:"taskName,omitempty"`
	Date string `json:"date,omitempty" bson:"date,omitempty"`
	Status bool `json:"status,omitempty" bson:"status,omitempty"`
}