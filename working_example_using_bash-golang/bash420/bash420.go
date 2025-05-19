package main

import (
	"fmt"
	"time"
)

func main() {
	exitPlease := 0
	inc := 0

	for exitPlease == 0 { // The loop will continue until exitPlease is not 0
		fmt.Printf("Boo %d\n", inc)
		inc++
		time.Sleep(1 * time.Second) // Sleep for 1 second
	}

	// Exit condition
	fmt.Println("Exiting...")
}

