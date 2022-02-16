package models
type Task struct {
	TaskName string `json:"taskName,omitempty" bson:"taskName,omitempty"`
	Date string `json:"date,omitempty" bson:"date,omitempty"`
	Status bool `json:"status,omitempty" bson:"status,omitempty"`
}