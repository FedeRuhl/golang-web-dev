package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// For this code to run, you will need this package:
// go get github.com/google/uuid

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.New()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true, // OVER HTTPS
			HttpOnly: true, // CAN'T ACCESS WITH JAVASCRIPT
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
