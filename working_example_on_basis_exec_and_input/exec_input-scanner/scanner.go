package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// getFormattedName returns the formatted full name.
func getFormattedName(firstName, lastName string) string {
	fullName := firstName + " " + lastName
	return strings.Title(fullName)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Println("\nPlease tell me your name: ")
		fmt.Println("(enter 'q' at any time to quit)")
		
		fmt.Print("First name: ")
		scanner.Scan()
		fName := scanner.Text()
		if strings.ToLower(fName) == "q" {
			break
		}
		
		fmt.Print("Last name: ")
		scanner.Scan()
		lName := scanner.Text()
		if strings.ToLower(lName) == "q" {
			break
		}
		
		formattedName := getFormattedName(fName, lName)
		fmt.Printf("\nHello, %s!\n", formattedName)
	}
}

// Please tell me your name: 
// (enter 'q' at any time to quit)
// First name: max
// Last name: ned

// Hello, Max Ned!

// Please tell me your name: 
// (enter 'q' at any time to quit)
// First name: q
