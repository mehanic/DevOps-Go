package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define the slices
	bicycles := []string{"trek", "cannondale", "redline", "specialized"}
	names := []string{"Tom", "Jordan", "Benjamin", "Jack", "Adam", "Dick", "Newton", "Darwin"}

	// Print the second to last bicycle in title case
	fmt.Println(strings.Title(bicycles[len(bicycles)-2]))

	// Print a message about the first bicycle
	message := "My first bicycle was a " + strings.Title(bicycles[0]) + "."
	fmt.Println(message)

	// Print a message about the favorite idol
	favoriteIdolMessage := "My favorite idol is " + names[len(names)-1] + "."
	fmt.Println(favoriteIdolMessage)
}

// Redline
// My first bicycle was a Trek.
// My favorite idol is Darwin.
