package models
type Task struct {
	TaskName string `json:"TaskName,omitempty" bson:"TaskName,omitempty"`
	Date string `json:"Date,omitempty" bson:"Date,omitempty"`
	Status string `json:"Status,omitempty" bson:"Status,omitempty"`
	User User `json:"User,omitempty" bson:"User,omitempty"`
}