package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to get a formatted full name (from the previous example)
func getFormattedName(first string, last string, middle ...string) string {
	var fullName string
	if len(middle) > 0 {
		fullName = first + " " + middle[0] + " " + last
	} else {
		fullName = first + " " + last
	}
	return strings.Title(fullName)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter 'q' at any time to quit.")
	for {
		fmt.Print("\nPlease give me a first name: ")
		scanner.Scan()
		first := scanner.Text()
		if first == "q" {
			break
		}

		fmt.Print("Please give me a last name: ")
		scanner.Scan()
		last := scanner.Text()
		if last == "q" {
			break
		}

		formattedName := getFormattedName(first, last)
		fmt.Println("\tNeatly formatted name: " + formattedName + ".")
	}
}

// Enter 'q' at any time to quit.

// Please give me a first name: taya
// Please give me a last name: naia
// 	Neatly formatted name: Taya Naia.
