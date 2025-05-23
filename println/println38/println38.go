package main

import (
	"fmt"
	"strings"
)

// Function to get a formatted full name
func getFormattedName(first string, last string, middle ...string) string {
	var fullName string

	// Check if a middle name was provided
	if len(middle) > 0 {
		fullName = first + " " + middle[0] + " " + last
	} else {
		fullName = first + " " + last
	}

	// Convert the full name to title case
	return strings.Title(fullName)
}

func main() {
	// Example usage
	nameWithMiddle := getFormattedName("john", "doe", "paul")
	fmt.Println(nameWithMiddle) // Output: John Paul Doe

	nameWithoutMiddle := getFormattedName("john", "doe")
	fmt.Println(nameWithoutMiddle) // Output: John Doe
}

// John Paul Doe
// John Doe
