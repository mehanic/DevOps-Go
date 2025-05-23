package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Prompt the user
	fmt.Print("Enter your name: ")

	// Read user input from stdin
	reader := bufio.NewReader(os.Stdin)
	myvar, _ := reader.ReadString('\n')

	// Remove the newline character from the input
	myvar = myvar[:len(myvar)-1]

	// Print the greeting
	fmt.Printf("Hello %s\n", myvar)
}

