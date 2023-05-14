package main

import (
	"html/template"
	"log"
	"os"
)

type menuItem struct {
	Name  string
	Price float64
}

type menuOption struct {
	Items []menuItem
}

type restaurantMenu struct {
	Breakfast, Lunch, Dinner menuOption
}

type restaurantMenues struct {
	RestaurantName string
	Menu           restaurantMenu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	menues := []restaurantMenues{
		{
			RestaurantName: "Restaurante 1",
			Menu: restaurantMenu{
				Breakfast: menuOption{
					Items: []menuItem{
						{
							Name:  "huevo revuelto",
							Price: 500,
						},
					},
				},
				Lunch: menuOption{
					Items: []menuItem{
						{
							Name:  "milanesa con papas fritas",
							Price: 1000,
						},
					},
				},
				Dinner: menuOption{
					Items: []menuItem{
						{
							Name:  "pizza a la piedra",
							Price: 1200,
						},
					},
				},
			},
		},
		{
			RestaurantName: "Restaurante 2",
			Menu: restaurantMenu{
				Breakfast: menuOption{
					Items: []menuItem{
						{
							Name:  "huevo revuelto",
							Price: 600,
						},
					},
				},
				Lunch: menuOption{
					Items: []menuItem{
						{
							Name:  "milanesa con papas fritas",
							Price: 1100,
						},
					},
				},
				Dinner: menuOption{
					Items: []menuItem{
						{
							Name:  "pizza a la piedra",
							Price: 1300,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, menues)
	if err != nil {
		log.Fatalln(err)
	}
}
