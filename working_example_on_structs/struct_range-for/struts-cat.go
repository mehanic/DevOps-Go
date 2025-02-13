package main

import "fmt"

// Cat struct represents a cat with a description and a name
type Cat struct {
	Desc string
	Name string
}

func main() {
	// Creating a slice of Cat structs
	cats := []Cat{
		{Desc: "chubby", Name: "Zophie"},
		{Desc: "fluffy", Name: "Pooka"},
	}

	// Loop through the cats slice and print each cat's information
	for _, cat := range cats {
		fmt.Printf("Name: %s, Description: %s\n", cat.Name, cat.Desc)
	}
}

// Name: Zophie, Description: chubby
// Name: Pooka, Description: fluffy
