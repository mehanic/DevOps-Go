package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt and read a single word
	fmt.Print("Hi, please type the word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word) // Trim the newline
	fmt.Println("The word you entered is:", word)

	// Prompt and read two words
	fmt.Print("Can you please enter two words? ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	words := strings.Split(input, " ")
	if len(words) >= 2 {
		fmt.Printf("Here is your input: \"%s\" \"%s\"\n", words[0], words[1])
	} else {
		fmt.Println("Please enter at least two words.")
	}

	// Prompt and read a response using default variable (similar to Bash's $REPLY)
	fmt.Print("How do you feel about Go programming? ")
	reply, _ := reader.ReadString('\n')
	reply = strings.TrimSpace(reply)
	fmt.Printf("You said \"%s\", I'm glad to hear that!\n", reply)

	// Prompt and read multiple colors into an array
	fmt.Print("What are your favorite colors? ")
	colorInput, _ := reader.ReadString('\n')
	colorInput = strings.TrimSpace(colorInput)
	colours := strings.Split(colorInput, " ")

	// Print favorite colors
	if len(colours) > 0 {
		fmt.Printf("My favorite colors are also %s, %s, and %s:-)\n", colours[0], colours[1], colours[2])
	} else {
		fmt.Println("Please enter your favorite colors.")
	}
}

