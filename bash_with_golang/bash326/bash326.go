package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a new reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	// Prompt for user input
	fmt.Print("Enter value: ")

	// Read the input from the user and store it in the 'value' variable
	value, _ := reader.ReadString('\n')

	// Print the entered value
	fmt.Println("You entered:", value)
}

