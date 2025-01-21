package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define the map with river names and their corresponding countries
	rivers := map[string]string{
		"Yangzijiang": "china",
		"nile":         "egypt",
		"amazon":       "southamerica",
	}

	// Iterate through the map and print the formatted output
	for river, country := range rivers {
		// Capitalize the river and country names
		riverTitle := strings.Title(river)
		countryTitle := strings.Title(country)

		// Print the result
		fmt.Printf("The %s runs through %s.\n", riverTitle, countryTitle)
	}
}

// The Yangzijiang runs through China.
// The Nile runs through Egypt.
// The Amazon runs through Southamerica.
