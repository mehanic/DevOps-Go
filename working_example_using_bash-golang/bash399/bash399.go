package main

import (
	"fmt"
)

func main() {
	// Index zero of VARIABLE is the char 'M' & is 14 bytes long
	variable := "My test string"
	// Slicing in Go: [start:end] gets the substring
	substring := variable[3:7] // Starts from index 3 and ends before index 7
	fmt.Println(substring)      // Output: "test"
}

