package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// List of pets
	myPets := []string{"Zophie", "Pooka", "Fat-tail"}

	// Ask user for a pet name
	fmt.Println("Enter a pet name:")

	// Reading input from user
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Remove newline and any extra spaces

	// Check if the name is in the list
	found := false
	for _, pet := range myPets {
		if pet == name {
			found = true
			break
		}
	}

	// Output the result
	if !found {
		fmt.Printf("I do not have a pet named %s\n", name)
	} else {
		fmt.Printf("%s is my pet.\n", name)
	}
}

// Enter a pet name:
// larry
// I do not have a pet named larry
