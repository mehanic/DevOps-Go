package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	heads := 0

	// Loop 1000 times
	for i := 1; i <= 1000; i++ {
		// Simulate a coin flip (0 or 1)
		if rand.Intn(2) == 1 {
			heads++
		}

		// Print halfway message
		if i == 500 {
			fmt.Println("Halfway done!")
		}
	}

	// Print the result
	fmt.Printf("Heads came up %d times.\n", heads)
}


// Halfway done!
// Heads came up 499 times.
