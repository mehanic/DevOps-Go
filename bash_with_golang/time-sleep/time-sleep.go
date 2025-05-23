package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Count: ")
	count := 0

	for {
		if count < 5 {
			count++
			// Move the cursor to the beginning of the line
			fmt.Printf("\rCount: %d", count)
			// Wait for 1 second
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	// Move the cursor to the next line after counting
	fmt.Println()
}

