package main

import (
	"golang-web-dev/042_mongodb/06_hands-on/starting-code/controllers"
	"golang-web-dev/042_mongodb/06_hands-on/starting-code/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[primitive.ObjectID]models.User {
	return map[primitive.ObjectID]models.User{}
}
