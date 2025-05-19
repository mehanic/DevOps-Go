package main

import (
	"fmt"
)

func main() {
	// Define the variable
	planet := "Earth"

	// Print the variable in different ways
	fmt.Println(planet)        // Prints the value of planet
	fmt.Println("$planet")     // Prints the literal string "$planet"
	fmt.Println(planet)        // Prints the value of planet again
	fmt.Println("$planet")     // Prints the dollar sign followed by 'planet' without escape

	// Prompt the user for input
	var userInput string
	fmt.Print("Enter some text: ")
	fmt.Scanln(&userInput) // Read user input

	// Display the updated variable
	fmt.Printf("$planet now equals %s\n", userInput) // Print the updated input
}

