package models

import "go.mongodb.org/mongo-driver/bson"

type User struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
	Id     bson.ObjectId `json:"id" bson:"_id"`
}

// Id was of type string before
