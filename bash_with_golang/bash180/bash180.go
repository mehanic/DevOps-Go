package main

import (
	"fmt"
)

func main() {
	var name string
	var correct string

	for correct != "y" {
		// Prompt for the user's name
		fmt.Print("Enter your name: ")
		fmt.Scanln(&name)

		// Ask for confirmation
		fmt.Printf("Is %s correct? ", name)
		fmt.Scanln(&correct)
	}
}

