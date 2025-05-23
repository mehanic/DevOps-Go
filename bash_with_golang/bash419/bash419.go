package main

import (
	"fmt"
	"time"
)

func main() {
	exitPlease := false
	inc := 0

	// Infinite loop simulating the until loop in Bash
	for !exitPlease { // Exit condition is never satisfied in the original script
		// Print the message, equivalent to the Bash echo command
		fmt.Printf("Boo %d\n", inc)

		// Increment the counter
		inc++

		// Sleep for 1 second
		time.Sleep(1 * time.Second)
	}

	// Note: The exit point won't be reached in this version
}

