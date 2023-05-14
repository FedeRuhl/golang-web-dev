package controllers

import (
	"encoding/json"
	"fmt"
	"golang-web-dev/042_mongodb/06_hands-on/starting-code/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	session map[primitive.ObjectID]models.User
}

func NewUserController(s map[primitive.ObjectID]models.User) *UserController {
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

	// Fetch user
	u, ok := uc.session[oid]

	if !ok {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = primitive.NewObjectID()

	// store the user in mongodb
	uc.session[u.Id] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !primitive.IsValidObjectID(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println(err)
	}

	// Delete user
	delete(uc.session, oid)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
