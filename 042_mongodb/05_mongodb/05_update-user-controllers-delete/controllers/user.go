package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-web-dev/042_mongodb/05_mongodb/05_update-user-controllers-delete/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	Client *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !primitive.IsValidObjectID(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation.
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println(err)
	}

	// composite literal
	u := models.User{}

	// Fetch user
	collection := uc.Client.Database("go-web-dev-db").Collection("users")
	err = collection.FindOne(context.Background(), bson.M{"_id": oid}).Decode(&u)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// var u models.User
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = primitive.NewObjectID()

	// store the user in mongodb
	collection := uc.Client.Database("go-web-dev-db").Collection("users")
	result, err := collection.InsertOne(context.Background(), u)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.InsertedID)
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !primitive.IsValidObjectID(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
	}

	// Delete user
	collection := uc.Client.Database("go-web-dev-db").Collection("users")
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": oid})

	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
