package main

import (
	"fmt"
)

func main() {
	VAR := 10 // Edit to 1 or 2 and re-run.

	switch VAR {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("What is this var?")
		return // Exit with a non-zero status.
	}
}

