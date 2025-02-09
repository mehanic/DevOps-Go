package main

import "fmt"

// Function to clear the names of the toys and return the count of toys picked up
func pickUp(toys []string) int {
	count := len(toys) // Get the number of toys
	for i := range toys {
		toys[i] = "" // Set each toy name to an empty string
	}
	return count // Return the number of toys picked up
}

func main() {
	// Initialize a slice of toy names
	toys := []string{"Winnie the Pooh", "ring tower", "wooden blocks"}
	
	// Call pickUp function to clear toy names
	count := pickUp(toys)
	
	// Print the result
	fmt.Printf("All done! Picked up %d toys.\n", count)
}

