package main

import (
	"html/template"
	"log"
	"os"
)

type Hotel struct {
	Name, Address, City, Zip, Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	californiaHotels := []Hotel{
		{
			"Hotel1",
			"Hotel1",
			"Hotel1",
			"Hotel1",
			"Southern",
		},
		{
			"Hotel2",
			"Hotel2",
			"Hotel2",
			"Hotel2",
			"Northern",
		},
	}

	err := tpl.Execute(os.Stdout, californiaHotels)
	if err != nil {
		log.Fatalln(err)
	}
}
