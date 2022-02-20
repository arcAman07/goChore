package config

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

// const uri = "mongodb://localhost:27017/taskDB"

func Configure() (*mongo.Collection, *mongo.Collection) {
	// Create a new client and connect to the server

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("MONGO_URI")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	UserCollection := client.Database("taskDB").Collection("users")
	TaskCollection := client.Database("taskDB").Collection("tasks")
	return UserCollection, TaskCollection
}
