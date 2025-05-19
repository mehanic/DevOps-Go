package main

import (
	"fmt"
)

func main() {
	// Declare and initialize the variable
	planet := "Earth"

	// Print the variable in different formats
	fmt.Println(planet)               // prints: Earth
	fmt.Println("$planet")            // prints: $planet
	fmt.Println(fmt.Sprintf("%s", planet)) // prints: Earth
	fmt.Println("$planet")            // prints: $planet

	// Prompt for user input
	fmt.Print("Enter some text: ")
	var userInput string
	fmt.Scanln(&userInput) // Read user input

	// Update the variable with user input
	planet = userInput

	// Print the updated variable
	fmt.Printf("$planet now equals %s\n", planet)

	// Exit the program
}

