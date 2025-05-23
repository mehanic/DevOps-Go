package main

import (
	"fmt"
)

func main() {
	var catNames []string // Create a slice to hold cat names
	var name string

	for {
		// Prompt for the cat name
		fmt.Printf("Enter the name of cat %d (Or enter nothing to stop):\n", len(catNames)+1)
		fmt.Scanln(&name)

		// Check if the user entered nothing
		if name == "" {
			break
		}

		// Append the name to the slice
		catNames = append(catNames, name)
	}

	// Display the cat names
	fmt.Println("The cat names are:")
	for _, catName := range catNames {
		fmt.Println("  " + catName)
	}
}

