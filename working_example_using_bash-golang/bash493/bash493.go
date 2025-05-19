package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Iterate over command-line arguments (words).
	for _, word := range os.Args[1:] {
		fmt.Println(word) // Print the word
		time.Sleep(1 * time.Second) // Sleep for 1 second
	}
}

