package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a new reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for input
	fmt.Print("Name few cities? ")

	// Read the input from the user
	input, _ := reader.ReadString('\n')

	// Trim whitespace and split the input into cities
	cities := strings.Fields(strings.TrimSpace(input))

	// Access and print the name of the city at index 2, if it exists
	if len(cities) > 2 {
		fmt.Printf("Name of city is %s.\n", cities[2])
	} else {
		fmt.Println("Not enough cities were provided.")
	}
}

