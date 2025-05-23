package main

import (
	"fmt"
)

func main() {
	ctr := 1 // Initialize the counter variable

	// Use a while-like loop with a for statement
	for ctr < 9 {
		fmt.Printf("CTR var: %d\n", ctr) // Print the counter value
		ctr++ // Increment the counter by 1
	}

	fmt.Println("Finished") // Print the final message
}

