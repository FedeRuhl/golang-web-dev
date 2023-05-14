package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/dog/", handleDog)
	http.HandleFunc("/me/", handleMe)
	http.ListenAndServe(":8080", nil)
}

func handleHome(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "home")
}

func handleDog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

func handleMe(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "It's me, Fede")
}
