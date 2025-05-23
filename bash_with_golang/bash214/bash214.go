package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start of program.")

	// Create a goroutine that sleeps for 5 seconds and prints a message.
	go takeANap()

	// The main function will not wait for the goroutine to complete
	fmt.Println("End of program.")

	// Sleep for a little to ensure the goroutine has time to finish.
	time.Sleep(6 * time.Second)
}

// takeANap simulates the nap with a 5-second sleep.
func takeANap() {
	time.Sleep(5 * time.Second)
	fmt.Println("Wake up!")
}

