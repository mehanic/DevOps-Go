package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	myPets := []string{"Zophie", "Pooka", "Fat-tail"}
	
	// Prompt for pet name
	fmt.Println("Enter a pet name:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Check if the pet name is in the list
	petExists := false
	for _, pet := range myPets {
		if pet == name {
			petExists = true
			break
		}
	}

	// Output the result
	if !petExists {
		fmt.Println("I do not have a pet named " + name)
	} else {
		fmt.Println(name + " is my pet.")
	}
}

