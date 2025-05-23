package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator to get different results each time
	rand.Seed(time.Now().UnixNano())

	// Loop to generate 5 random numbers
	for i := 0; i < 5; i++ {
		// Generate a random integer between 1 and 10
		fmt.Println(rand.Intn(10) + 1)
	}
}

