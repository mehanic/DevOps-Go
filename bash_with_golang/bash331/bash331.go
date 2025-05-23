package main

import (
	"fmt"
	"os"
)

func main() {
	// Print command and arguments
	fmt.Println("  Command is", os.Args[0])
	if len(os.Args) > 1 {
		fmt.Println("  First Argument is", os.Args[1])
	} else {
		fmt.Println("  First Argument is not provided")
	}
	if len(os.Args) > 2 {
		fmt.Println("  Second Argument is", os.Args[2])
	} else {
		fmt.Println("  Second Argument is not provided")
	}
}

