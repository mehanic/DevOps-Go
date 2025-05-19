package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the name was provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a name as an argument.")
		return
	}

	// Assign the first argument to the name variable
	name := os.Args[1]

	// Print the story using the provided name
	fmt.Printf("There once was a person named %s. %s enjoyed Linux and Bash so much that he/she wrote a book about it! %s really hopes everyone enjoys his/her book.\n", name, name, name)
}

