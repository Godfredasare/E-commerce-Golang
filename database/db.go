package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

var dbName = "Tonaton"

func Collection(name string) *mongo.Collection {
	if client == nil {
		log.Fatal("Database client is not initialized. Call InitDB first.")
	}
	return client.Database(dbName).Collection(name)
}

func InitDB() {
	var err error

	db_UUrl := os.Getenv("MONGODB_URI")

	options := options.Client().ApplyURI(db_UUrl)

	client, err = mongo.Connect(context.Background(), options)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("DB connected sucsessfully")
	createUniqueIndex()
}

func CloseDB() {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
}

func createUniqueIndex() {
	col := Collection("users")
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, // Create an index on the email field
		Options: options.Index().SetUnique(true),
	}

	_, err := col.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Could not create index: %v", err)
	}

	fmt.Println("Unique index on email field created successfully")
}
