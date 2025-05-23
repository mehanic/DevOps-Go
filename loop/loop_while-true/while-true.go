package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Function to get a formatted full name with proper Unicode title-casing
func getFormattedName(first string, last string, middle ...string) string {
	var fullName string
	if len(middle) > 0 {
		fullName = first + " " + middle[0] + " " + last
	} else {
		fullName = first + " " + last
	}

	// Используем cases.Title для правильного форматирования заглавных букв
	titleCaser := cases.Title(language.Und)
	return titleCaser.String(fullName)
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

		fmt.Print("Please give me a last name: ")
		scanner.Scan()
		middle := scanner.Text()
		if last == "q" {
			break
		}

		formattedName := getFormattedName(first, last, middle)
		fmt.Println("\tNeatly formatted name: " + formattedName + ".")
	}
}


// go run while-true.go 
// Enter 'q' at any time to quit.

// Please give me a first name: mike
// Please give me a last name: london
// Please give me a last name: deiry
// 	Neatly formatted name: Mike Deiry London.

// Please give me a first name: q
