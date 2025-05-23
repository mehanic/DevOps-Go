package main

import (
	"fmt"
	"time"
)

func main() {
	x := 10
	count := 0 // To limit the number of iterations

	// Using a for loop to replicate the while loop behavior
	for x == 10 && count < 5 {
		fmt.Println(x)
		time.Sleep(2 * time.Second) // Sleep for 2 seconds
		count++                      // Increment count
	}
}

