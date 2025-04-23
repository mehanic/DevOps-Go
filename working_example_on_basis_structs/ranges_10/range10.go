package main

import (
	"fmt"
//	"sort"
	"strings"
)

func main() {
	// Define a slice of pairs to preserve the order of insertion
	favoriteLanguages := []struct {
		name     string
		language string
	}{
		{"jen", "python"},
		{"sarah", "c"},
		{"edward", "ruby"},
		{"phi", "python"},
	}

	// Iterating over the favoriteLanguages slice and printing each person's favorite language
	for _, entry := range favoriteLanguages {
		fmt.Printf("%s's favorite language is %s.\n", strings.Title(entry.name), strings.Title(entry.language))
	}
}


// Jen's favorite language is Python.
// Sarah's favorite language is C.
// Edward's favorite language is Ruby.
// Phi's favorite language is Python.
