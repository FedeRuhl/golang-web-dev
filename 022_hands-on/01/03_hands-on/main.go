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
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/dog/", handleDog)
	http.HandleFunc("/me/", handleMe)
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
