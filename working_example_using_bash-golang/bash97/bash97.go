package main

import (
	"fmt"
)

func main() {
	// Simulate the loop from 1 to 6
	for i := 1; i <= 6; i++ {
		// Set -x equivalent: Start debug
		fmt.Printf("Executing command: echo %d\n", i)
		
		// Print the current number
		fmt.Println(i)

		// Set +x equivalent: End debug
	}
	
	// Final message
	fmt.Println("finish.")
}

