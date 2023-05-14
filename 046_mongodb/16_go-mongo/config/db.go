package config

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// database
var DB *mongo.Database

// collections
var Books *mongo.Collection

func init() {
	// get a mongo sessions
	//s, err := mgo.Dial("mongodb://bond:moneypenny007@localhost/bookstore")

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://bond:moneypenny007@localhost:27017/?ssl=false")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)

	// Check if connection error, is MongoDB running?
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	DB = client.Database("bookstore")
	Books = DB.Collection("books")

	fmt.Println("You connected to your mongo database.")
}
