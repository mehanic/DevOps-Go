package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Print the script name and command-line arguments
	fmt.Printf("Executing script %s\n", os.Args[0])

	if len(os.Args) > 3 {
		fmt.Println(os.Args[1], os.Args[2], os.Args[3])
	} else {
		fmt.Println("Not enough arguments provided")
	}

	// Simulating the 'set' behavior with values 'eins zwei drei'
	germanNumbers := []string{"eins", "zwei", "drei"}
	fmt.Println("One two three in German are:")
	for _, num := range germanNumbers {
		fmt.Println(num)
	}

	// Simulating 'set' with the textline "name phone address birthdate salary"
	textline := "name phone address birthdate salary"
	words := strings.Fields(textline) // Split textline into words by spaces
	fmt.Println(strings.Join(words, " "))

	// Access individual elements like $1 and $4 in Bash
	if len(words) > 3 {
		fmt.Printf("At this time $1 = %s and $4 = %s\n", words[0], words[3])
	} else {
		fmt.Println("Not enough words to access $1 and $4")
	}
}

