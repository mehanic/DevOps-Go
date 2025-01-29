package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define the slice with names
	names := []string{"chidi", "peter", "ada", "feldor", "mickael"}

	// Print each name with different string manipulations
	fmt.Println(strings.ToLower(names[0])) // Convert all letters to lowercase
	fmt.Println(strings.Title(names[1]))   // Capitalize the first letter of each word
	fmt.Println(strings.ToUpper(names[2])) // Convert all letters to uppercase
	fmt.Println("\t" + names[3] + "\n")    // Print with tab and newline
	fmt.Println(names[4])                  // Print the name as is
}

