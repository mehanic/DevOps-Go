package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	// Try to read the /etc/shadow file
	_, err := ioutil.ReadFile("/etc/shadow")
	if err != nil {
		// Print the error message and exit with code 123
		fmt.Println("Error reading /etc/shadow:", err)
		os.Exit(123)
	}

	// If successful, we could process the file here (not done since we expect it to fail)
	fmt.Println("Successfully read /etc/shadow")
}

