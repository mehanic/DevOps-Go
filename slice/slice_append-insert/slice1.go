package main

import (
	"fmt"
)

func main() {
	// Initialize a slice of guests
	guests := []string{"chima", "daniel", "kosomo"}

	// Insert "jennifer" at the beginning (index 0)
	guests = append([]string{"jennifer"}, guests...)

	// Insert "nnamdi" at index 2
	index := 2
	guests = append(guests[:index+1], guests[index:]...)
	guests[index] = "nnamdi"

	// Append "kim" to the end of the slice
	guests = append(guests, "kim")

	// Print the number of guests and the list of guests
	fmt.Printf("These %d people have been invited:\n", len(guests))
	fmt.Println(guests)
}

