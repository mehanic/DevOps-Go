package main

import (
	"fmt"
	"os"
)

func main() {
	count := len(os.Args) - 1
	fmt.Printf("There are %d names.\n", count)
}

// main.go Alice Bob Charlie
// There are 3 names.
