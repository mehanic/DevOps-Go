package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a directory.
	err := os.Mkdir("/tmp/temp_dir", 0755) // 0755 is the directory permission
	var mkdirRC int
	if err != nil {
		mkdirRC = 1 // Set return code to 1 if mkdir fails
		fmt.Println("Error creating directory:", err)
	} else {
		mkdirRC = 0 // Set return code to 0 if mkdir succeeds
	}

	// Use os.Stat to check if the directory was created.
	_, err = os.Stat("/tmp/temp_dir")
	var testRC int
	if os.IsNotExist(err) {
		testRC = 1 // Directory does not exist
	} else {
		testRC = 0 // Directory exists
	}

	// Check out the return codes:
	fmt.Printf("mkdir resulted in %d, test resulted in %d.\n", mkdirRC, testRC)
}

