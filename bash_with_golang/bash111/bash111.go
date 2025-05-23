package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator to get different results each time
	rand.Seed(time.Now().UnixNano())

	// Loop from 1 to 100
	for i := 1; i <= 100; i++ {
		// Generate a random number
		randomNumber := rand.Intn(10000) // Change the range as per your requirement
		fmt.Printf("%d: %d\n", i, randomNumber)
	}
}

