package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	prompt := "\nPlease enter the name of a city you have visited:\n"
	prompt += "\nEnter 'quit' when you are finished. "

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		// Scan user input
		scanner.Scan()
		city := scanner.Text()

		// Check for "quit" condition
		if city == "quit" {
			break
		} else {
			// Capitalize the city name and print the message
			fmt.Println("I'd love to go to " + strings.Title(city) + "!")
		}
	}
}

