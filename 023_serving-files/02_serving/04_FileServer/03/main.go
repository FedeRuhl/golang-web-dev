package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	// SERVE "./assets" FOLDER
	// BUT SINCE THERE IS NO "/assets" INSIDE "./assets" FOLDER
	// REMOVE "/assets" FROM THE URL

	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", dog)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// IF "/assets" EXISTS INSIDE "./assets" FOLDER
	// YOU CAN DO THIS:
	// http.Handle("/assets/", http.FileServer(http.Dir("assets")))
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/assets/toby.jpg">`)
}
