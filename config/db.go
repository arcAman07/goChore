package config

import (
	"fmt"

	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

// database
var DB *mgo.Database

// collections
var Books *mgo.Collection

func init() {
	// get a mongo sessions
	//s, err := mgo.Dial("mongodb://bond:moneypenny007@localhost/bookstore")
	s, err := mgo.Dial("mongodb://localhost:27017/tasksDB")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("tasksDB")
	Books = DB.C("tasks")

	fmt.Println("You connected to your mongo database.")
}
