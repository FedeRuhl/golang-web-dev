package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(handleHome))
	http.HandleFunc("/dog/", http.HandlerFunc(handleDog))
	http.HandleFunc("/me/", http.HandlerFunc(handleMe))
	http.ListenAndServe(":8080", nil)
}

func handleHome(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "home.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func handleDog(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func handleMe(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "me.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
