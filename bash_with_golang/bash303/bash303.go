package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Define the file name to check
	fileName := "myfile"

	// Loop until the file exists
	for {
		// Check if the file exists
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			// File does not exist, sleep for 1 second
			time.Sleep(1 * time.Second)
		} else {
			// File exists, break the loop
			fmt.Println("File exists!")
			break
		}
	}
}

