package main

import (
	"golang-web-dev/042_mongodb/10_hands-on/starting-code/controllers"
	"net/http"
)

func main() {
	s := controllers.NewSessionController()
	http.HandleFunc("/", s.Index)
	http.HandleFunc("/bar", s.Bar)
	http.HandleFunc("/signup", s.Signup)
	http.HandleFunc("/login", s.Login)
	http.HandleFunc("/logout", s.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
