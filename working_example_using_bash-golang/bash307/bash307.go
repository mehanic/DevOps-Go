package main

import (
	"fmt"
)

func main() {
	// Initialize the counter
	count := 0

	// Go's equivalent of an until loop
	for count <= 5 {
		fmt.Printf("Value of count is: %d\n", count)
		count++ // Increment the count
	}
}

