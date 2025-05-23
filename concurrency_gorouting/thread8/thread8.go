package main

import (
	"fmt"
	"time"
)

func takeANap() {
	time.Sleep(5 * time.Second)
	fmt.Println("Wake up!")
}

func main() {
	fmt.Println("Start of program.")

	// Launch the takeANap function as a goroutine
	go takeANap()

	fmt.Println("End of program.")

	// Wait for a moment to allow the goroutine to finish
	// This is to ensure that the program doesn't exit immediately
	time.Sleep(6 * time.Second) // Sleep for longer than the goroutine's sleep time
}

// ╰─λ go run thread8.go                                                                                                            0 (0.001s) < 14:41:20
// Start of program.
// End of program.
// Wake up!
