package main

import (
	"fmt"
)

func main() {
	// Initialize the counter
	count := 6

	// Go's equivalent of a while loop
	for count > 0 {
		fmt.Printf("Value of count is: %d\n", count)
		count-- // Decrease the count
	}
}

