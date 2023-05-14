package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	} else {
		collection := client.Database("tests-db").Collection("tests")
		cursor, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			fmt.Println("Error getting collection:", err)
			return
		}
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var doc bson.M
			if err := cursor.Decode(&doc); err != nil {
				fmt.Println("Error decoding document:", err)
				return
			}
			fmt.Println(doc)
		}
	}
}
