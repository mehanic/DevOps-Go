package main

import (
	"fmt"
	"os"
)

func main() {
	// Capture the program name (equivalent to $0 in Bash)
	ourFilename := os.Args[0]
	fmt.Println(ourFilename) // Print the program name

	// Capture the number of arguments (equivalent to $# in Bash)
	numArguments := len(os.Args) - 1 // Exclude the program name from the count
	fmt.Println("The total number of arguments is:", numArguments)

	// Print the first three arguments if they exist
	if numArguments > 0 {
		fmt.Println("The first argument is:", os.Args[1])
	}
	if numArguments > 1 {
		fmt.Println("The second argument is:", os.Args[2])
	}
	if numArguments > 2 {
		fmt.Println("The third argument is:", os.Args[3])
	}
}

