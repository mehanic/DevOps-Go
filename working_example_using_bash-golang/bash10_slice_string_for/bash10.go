package main

import (
	"fmt"
)

func main() {
	// Declare a slice of strings to hold file names
	files := []string{"file1", "file2", "file3"}

	// Iterate over each element in the slice
	for _, element := range files {
		fmt.Println(element) // Print each file name
	}

	// Print the final message
	fmt.Println("Echo'd all the files")
}

