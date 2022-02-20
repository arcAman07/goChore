package models

// LoggedIn in the schema tells us whether the user is logged in or not

// 0 -> loggedOut, 1 -> loggedIn

type User struct {
	Username       string `json:"Username,omitempty" bson:"Username,omitempty"`
	Password       string `json:"Password,omitempty" bson:"Password,omitempty"`
	HashedPassword string `json:"HashedPassword,omitempty" bson:"HashedPassword,omitempty"`
	LoggedIn       string `json:"LoggedIn,omitempty" bson:"LoggedIn,omitempty"`
}
