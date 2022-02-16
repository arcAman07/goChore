package models
type User struct {
	ID int64 `json:"id,omitempty" bson:"id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	HashedPassword string `json:"hashed_password,omitempty" bson:"hashed_password,omitempty"`
}