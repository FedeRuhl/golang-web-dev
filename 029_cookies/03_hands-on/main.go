package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("counter")
	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:  "counter",
			Value: "1",
			Path:  "/",
		})
	} else {
		old, err := strconv.Atoi(c.Value)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		c.Value = strconv.Itoa(old + 1)
		http.SetCookie(w, c)
	}

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("counter")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "COUNTER:", c)
}

// Using cookies, track how many times a user has been to your website domain.
