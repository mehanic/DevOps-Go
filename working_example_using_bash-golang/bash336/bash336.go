package main

import (
	"fmt"
//	"os"
)

func main() {
	// Declare a filename variable
	filename := "test1"

	// Print a multi-line string, similar to a here-document in Bash
	fmt.Println(`When we add quotes before and after here
Document marker, we can include variables
Such as $USER, $PATH, ` + filename + ` and similar`)
}

