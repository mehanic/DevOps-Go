package main

import (
	"fmt"
	"os"
)

// Function to reverse a string.
func reverser(input string) string {
	runes := []rune(input)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	// Check if exactly one argument is passed.
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments!")
		fmt.Println("Usage: ./reverser <input-to-be-reversed>")
		os.Exit(1)
	}

	// Capture the user input and wrap it with underscores for readability.
	userInput := "_" + os.Args[1] + "_"

	// Reverse the input string using the reverser function.
	reversedInput := reverser(userInput)

	// Output the reversed string.
	fmt.Println("Your reversed input is:", reversedInput)
}

