package main

import (
	"fmt"
)

func main() {
	// Initialize CTR variable
	CTR := 1

	// While loop equivalent
	for CTR < 9 {
		fmt.Printf("CTR var: %d\n", CTR)
		CTR++ // Increment CTR by 1
	}

	// Print finished message
	fmt.Println("Finished")
}

