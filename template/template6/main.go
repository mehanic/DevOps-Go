package main

import (
	"html/template"
	"net/http"
)

var baseTemplate *template.Template

type ViewData struct {
	User User
}

type User struct {
	ID    int
	Email string
}

func main() {
	var err error

	// Parse the base template with no user-specific logic
	baseTemplate, err = template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)

	println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:    1,
		Email: "jon@calhoun.io",
	}

	vd := ViewData{User: user}

	// Clone the template per request to avoid race conditions
	clonedTmpl, err := baseTemplate.Clone()
	if err != nil {
		http.Error(w, "Template clone error", http.StatusInternalServerError)
		return
	}

	// Set up the user-specific FuncMap
	clonedTmpl = clonedTmpl.Funcs(template.FuncMap{
		"hasPermission": func(feature string) bool {
			// Your custom permission logic
			return user.ID == 1 && feature == "feature-a"
		},
	})

	// Set the content type
	w.Header().Set("Content-Type", "text/html")

	// Execute the safe cloned template
	err = clonedTmpl.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
