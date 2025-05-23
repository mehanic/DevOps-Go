package main

import (
	"fmt"
	"os"
)

func main() {
	// Redirecting STDOUT to STDERR
	fmt.Fprintln(os.Stderr, "Redirect this STDOUT to STDERR")
}

