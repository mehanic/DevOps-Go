package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the hostname
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
		return
	}

	// Print the message with the hostname
	fmt.Printf("This script is running on %s.\n", hostName)
}

