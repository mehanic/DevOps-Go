package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define favoriteLanguages map with mixed types (interface{})
	favoriteLanguages := map[string]interface{}{
		"jen":    []string{"python", "ruby"},
		"sarah":  "c",
		"edward": []string{"ruby", "go"},
		"phil":   []string{"python", "haskell"},
		"Jack":   "c",
	}

	// Iterate through the favoriteLanguages map
	for name, languages := range favoriteLanguages {
		fmt.Printf("\n%s's favorite languages are:\n", strings.Title(name))

		// Check if the value is a string or a slice of strings
		switch langs := languages.(type) {
		case string:
			fmt.Printf("\t%s\n", strings.Title(langs))
		case []string:
			for _, language := range langs {
				fmt.Printf("\t%s\n", strings.Title(language))
			}
		}
	}
}

// Edward's favorite languages are:
// 	Ruby
// 	Go

// Phil's favorite languages are:
// 	Python
// 	Haskell

// Jack's favorite languages are:
// 	C

// Jen's favorite languages are:
// 	Python
// 	Ruby

// Sarah's favorite languages are:
// 	C
