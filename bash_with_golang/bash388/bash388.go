package main

import (
	"fmt"
)

func main() {
	// Define a slice with file names
	files := []string{"file1", "file2", "file3"}

	// Iterate over the slice and print each element
	for _, element := range files {
		fmt.Println(element)
	}

	// Print the final message
	fmt.Println("Echo'd all the files")
}

