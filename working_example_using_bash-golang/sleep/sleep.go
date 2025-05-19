package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	count := 0
	fmt.Print("Count: ") // Print "Count:" without newline

	// Create a ticker that ticks every second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Handle interrupts (Ctrl+C) to exit gracefully
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	for {
		select {
		case <-ticker.C:
			if count < 40 {
				count++
				fmt.Printf("\rCount: %d", count) // Print on the same line with "\r"
			} else {
				fmt.Println("\nDone!")
				return
			}
		case <-quit:
			fmt.Println("\nInterrupted!")
			return
		}
	}
}
