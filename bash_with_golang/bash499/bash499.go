package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Infinite loop, only exits on the correct answer.
	for {
		fmt.Print("I have keys but no locks. I have a space but no room. You can enter, but can’t go outside. What am I? ")
		answer, _ := reader.ReadString('\n') // Read input until newline

		// Compile regex to match "keyboard" or "Keyboard"
		matched, err := regexp.MatchString(`(?i)keyboard`, answer) // Case insensitive match
		if err != nil {
			fmt.Println("Error in regex:", err)
			return
		}

		if matched {
			fmt.Println("Correct, congratulations!")
			break // Exit the loop.
		} else {
			// Print an error message and go back into the loop.
			fmt.Println("Incorrect, please try again.")
		}
	}

	// This will run after the break in the while loop.
	fmt.Println("Now we can continue after the while loop is done, awesome!")
}

