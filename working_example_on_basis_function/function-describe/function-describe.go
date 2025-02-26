package main

import (
	"fmt"
	"strings"
)

// describePet displays information about pets
func describePet(petName string, animalType ...string) {
	// Handle default value for animalType
	animal := "dog" // default value
	if len(animalType) > 0 {
		animal = animalType[0]
	}

	// Display pet information
	fmt.Printf("\nI have a %s.\n", animal)
	fmt.Printf("My %s's name is %s.\n", animal, strings.Title(petName))
}

func main() {
	// Calling the function with one argument
	describePet("william")
}

