package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for a word
	fmt.Print("Hi, please type the word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)
	fmt.Printf("The word you entered is: %s\n", word)

	// Prompt the user for favorite colors
	fmt.Print("What are your favorite colors? (Please separate by space): ")
	colorsInput, _ := reader.ReadString('\n')
	colors := strings.Fields(colorsInput) // Splits by space

	// Output the first three favorite colors, if available
	if len(colors) >= 3 {
		fmt.Printf("My favorite colors are: %s %s %s :)\n", colors[0], colors[1], colors[2])
	} else {
		fmt.Println("You didn't provide at least three colors!")
	}
}

