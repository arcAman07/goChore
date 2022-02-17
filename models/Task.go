package models
type Task struct {
	TaskName string `json:"TaskName,omitempty" bson:"TaskName,omitempty"`
	Date string `json:"Date,omitempty" bson:"Date,omitempty"`
	Status int `json:"Status,omitempty" bson:"Status,omitempty"`
	UserName string `json:"UserName,omitempty" bson:"UserName,omitempty"`
}

// Status is used to keep track of the status of the task

// 0 -> task is pending, 1 -> task is completed