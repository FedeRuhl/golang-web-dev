package main

import (
	"encoding/json"
	"golang-web-dev/042_mongodb/08_hands-on/07_solution/controllers"
	"golang-web-dev/042_mongodb/08_hands-on/07_solution/models"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
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

func getSession() map[string]models.User {
	m := make(map[string]models.User)

	content, err := os.ReadFile("db.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &m)

	if err != nil {
		log.Fatal(err)
	}

	return m
}
