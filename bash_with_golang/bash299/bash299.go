package main

import (
	"fmt"
	"os"
)

func main() {
	directory := "./BashScripting"

	// Check if the directory exists
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		fmt.Println("Directory does not exist")
	} else {
		fmt.Println("Directory exists")
	}
}

