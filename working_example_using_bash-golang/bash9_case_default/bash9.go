package main

import (
	"fmt"
)

func main() {
	variable := 10 // Edit to 1 or 2 and re-run

	switch variable {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("What is this var?")
		return // Exiting the program with a non-zero status can be done with os.Exit(1)
	}
}

