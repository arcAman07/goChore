package models

// LoggedIn in the schema tells us whether the user is logged in or not

// 0 -> loggedOut, 1 -> loggedIn

type User struct {
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	HashedPassword string `json:"hashed_password,omitempty" bson:"hashed_password,omitempty"`
	LoggedIn int `json:"logged_in,omitempty" bson:"logged_in,omitempty"`
}