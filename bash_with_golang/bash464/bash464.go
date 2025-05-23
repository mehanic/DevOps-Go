package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if the arguments are provided, otherwise prompt the user.
	characterName := getArgumentOrPrompt(1, "Name a fictional character: ")
	location := getArgumentOrPrompt(2, "Name an actual location: ")
	food := getArgumentOrPrompt(3, "What's your favorite food? ")

	// Compose and display the story.
	fmt.Printf("Recently, %s was seen in %s eating %s!\n", characterName, location, food)
}

// getArgumentOrPrompt checks if the argument is provided, otherwise prompts the user.
func getArgumentOrPrompt(argIndex int, prompt string) string {
	if len(os.Args) > argIndex {
		return os.Args[argIndex]
	}
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

