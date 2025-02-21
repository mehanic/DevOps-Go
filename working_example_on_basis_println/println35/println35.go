package main

import (
	"fmt"
	"strings"
)

func main() {
	// List of my favorite cars (Go slice)
	cars := []string{"mercedes", "ford", "dodge", "chevrolet", "rolls royce", "tesla"}

	// Printing messages using items from the slice
	fmt.Printf("I love %s cars.\n", strings.Title(cars[len(cars)-1])) // cars[-1] in Python
	fmt.Printf("I also love a few from %s, %s and %s.\n", 
		strings.Title(cars[len(cars)-2]), 
		strings.Title(cars[len(cars)-3]), 
		strings.Title(cars[len(cars)-4]))
	fmt.Printf("But my favourites are designed by %s and %s.\n", 
		strings.Title(cars[0]), 
		strings.Title(cars[1]))
}

// I love Tesla cars.
// I also love a few from Rolls Royce, Chevrolet and Dodge.
// But my favourites are designed by Mercedes and Ford.
