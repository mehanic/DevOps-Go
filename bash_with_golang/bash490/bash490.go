package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Generate a random number from 1-10.
	randomNumber := rand.Intn(10) + 1

	// Iterate over all possible random numbers.
	for number := 1; number <= 10; number++ {
		if number == randomNumber {
			fmt.Printf("Random number found: %d.\n", number)
			break // As soon as we have found the number, stop.
		}

		// If we get here, the number did not match.
		fmt.Printf("Number does not match: %d.\n", number)
	}

	fmt.Println("Number has been found, all done.")
}

