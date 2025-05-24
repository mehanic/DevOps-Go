package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// Define the map with river names and their corresponding countries
	rivers := map[string]string{
		"yangzijiang": "china",
		"nile":        "egypt",
		"amazon":      "southamerica",
	}

	// Create a caser for title case
	caser := cases.Title(language.English)

	// Iterate through the map and print the formatted output
	for river, country := range rivers {
		riverTitle := caser.String(river)
		countryTitle := caser.String(country)

		fmt.Printf("The %s runs through %s.\n", riverTitle, countryTitle)
	}
}

// avg function is called a variadic function since it take a slice of float64 values

// [12 23 45 45]
// []float64
// 31.25
