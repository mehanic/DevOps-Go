package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

type Widget struct {
	Name  string
	Price int
}

type ViewData struct {
	Name    string
	Widgets []Widget
}

func main() {
	var err error

	// Parse the single .gohtml template file with defined sections
	tmpl, err = template.ParseFiles("context.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)

	println("Server is running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := ViewData{
		Name: "John Smith",
		Widgets: []Widget{
			{"Blue Widget", 12},
			{"Red Widget", 15},
			{"Green Widget", 10},
		},
	}
	err := tmpl.ExecuteTemplate(w, "main", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
