package main

import (
	"fmt"
)

func main() {
	// Look at numbers 1-20, in steps of 2.
	for number := 1; number <= 20; number += 2 {
		if number%5 == 0 {
			continue // Unlucky number, skip this!
		}

		// Show the user which number we've processed.
		fmt.Printf("Looking at number: %d.\n", number)
	}
}

