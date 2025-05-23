package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define the map with people's names and their favorite languages
	favoriteLanguages := map[string]string{
		"jen":    "python",
		"sarah":  "c",
		"edward": "ruby",
		"phil":   "python",
		"jack":   "c",
	}

	// Define the slice of investigated people
	investigatedPeople := []string{"jen", "jack"}

	// Convert the slice to a map for O(1) lookup time
	investigatedMap := make(map[string]bool)
	for _, name := range investigatedPeople {
		investigatedMap[name] = true
	}

	// Iterate through the map and print the appropriate messages
	for name := range favoriteLanguages {
		fmt.Printf("Hi %s\n", strings.Title(name))
		if _, found := investigatedMap[name]; found {
			fmt.Println("Thanks for your investigation!")
		} else {
			fmt.Println("Would you take part in our investigation?")
		}
	}
}

// Hi Jen
// Thanks for your investigation!

// Hi Sarah
// Would you take part in our investigation?

// Hi Edward
// Would you take part in our investigation?

// Hi Phil
// Would you take part in our investigation?

// Hi Jack
// Thanks for your investigation!

