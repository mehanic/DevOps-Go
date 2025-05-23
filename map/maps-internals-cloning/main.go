package main

import (
	"fmt"
)

func main() {
	// Define example queries for testing
	testQueries := []string{"good", "iyi", "perfect", "harika", "notfound"}

	// Dictionary for English to Turkish
	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel",
	}

	// Remove specific entries for demonstration purposes
	delete(dict, "awesome")

	// Create reverse dictionary for Turkish to English
	turkish := make(map[string]string, len(dict))
	for k, v := range dict {
		turkish[v] = k
	}

	// Process each query in testQueries
	for _, query := range testQueries {
		fmt.Printf("Query: %q\n", query)

		// Lookup in English-to-Turkish dictionary
		if value, ok := dict[query]; ok {
			fmt.Printf("English: %q -> Turkish: %q\n\n", query, value)
			continue
		}

		// Lookup in Turkish-to-English dictionary
		if value, ok := turkish[query]; ok {
			fmt.Printf("Turkish: %q -> English: %q\n\n", query, value)
			continue
		}

		// If not found in either dictionary
		fmt.Printf("Sorry, %q not found.\n\n", query)
	}
}

// Query: "good"
// English: "good" -> Turkish: "iyi"

// Query: "iyi"
// Turkish: "iyi" -> English: "good"

// Query: "perfect"
// English: "perfect" -> Turkish: "mükemmel"

// Query: "harika"
// Turkish: "harika" -> English: "great"

// Query: "notfound"
// Sorry, "notfound" not found.
