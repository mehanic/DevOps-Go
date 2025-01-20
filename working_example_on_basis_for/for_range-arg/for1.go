package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	responses := make(map[string]string)
	pollingActive := true

	scanner := bufio.NewScanner(os.Stdin)

	for pollingActive {
		// Get user name
		fmt.Print("\nWhat is your name? ")
		scanner.Scan()
		name := scanner.Text()

		// Get user's response
		fmt.Print("Which mountain would you like to climb someday? ")
		scanner.Scan()
		response := scanner.Text()

		// Store the response
		responses[name] = response

		// Check if another person wants to respond
		fmt.Print("Would you like to let another person respond? (yes/no) ")
		scanner.Scan()
		repeat := scanner.Text()

		if strings.ToLower(repeat) == "no" {
			pollingActive = false
		}
	}

	fmt.Println("\n--- Poll Results ---")
	for name, response := range responses {
		fmt.Printf("%s would like to climb %s.\n", name, response)
	}
}

