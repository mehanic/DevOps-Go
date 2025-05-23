package main

import (
	"fmt"
	"os"
)

func main() {
	var testRc, simpleRc, extendedRc int

	// Test if the /tmp/ directory exists using the full command
	_, err := os.Stat("/tmp/")
	if err == nil {
		testRc = 0 // Directory exists
	} else if os.IsNotExist(err) {
		testRc = 1 // Directory does not exist
	} else {
		testRc = 1 // Some other error occurred
	}

	// Test if the /tmp/ directory exists using the simple shorthand
	_, err = os.Stat("/tmp/")
	if err == nil {
		simpleRc = 0 // Directory exists
	} else if os.IsNotExist(err) {
		simpleRc = 1 // Directory does not exist
	} else {
		simpleRc = 1 // Some other error occurred
	}

	// Test if the /tmp/ directory exists using the extended shorthand
	_, err = os.Stat("/tmp/")
	if err == nil {
		extendedRc = 0 // Directory exists
	} else if os.IsNotExist(err) {
		extendedRc = 1 // Directory does not exist
	} else {
		extendedRc = 1 // Some other error occurred
	}

	// Print the results
	fmt.Printf("The return codes are: %d, %d, %d.\n", testRc, simpleRc, extendedRc)
}

