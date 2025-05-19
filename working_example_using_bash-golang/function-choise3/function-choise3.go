package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// choose function that mimics the Bash choose function
func choose(defaultChoice, prompt string, yesAction func(), noAction func()) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)

	// Read input from the user
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	// If the answer is empty, use the default choice
	if answer == "" {
		answer = defaultChoice
	}

	// Execute actions based on the user's choice
	switch strings.ToLower(answer) {
	case "y", "1":
		yesAction()
	case "n", "0":
		noAction()
	default:
		fmt.Fprintf(os.Stderr, "Unexpected answer '%s'!\n", answer)
	}
}

// Example yes action
func yesAction() {
	fmt.Println("You chose to play a game!")
	// Here you can call any external program or perform actions
}

// Example no action
func noAction() {
	fmt.Println("See you later Professor Falkin.")
}

func main() {
	// Example usage of the choose function
	choose("y", "Do you want to play a game? ", yesAction, noAction)
}

