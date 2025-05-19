package main

import (
	"fmt"
	"time"
)

func main() {
	// Infinite loop.
	for {
		fmt.Println("Hello!")
		time.Sleep(1 * time.Second) // Wait for 1 second.
	}
}

