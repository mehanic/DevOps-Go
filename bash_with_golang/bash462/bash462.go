package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a new reader to capture input
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for information and read input
	fmt.Print("Name a fictional character: ")
	characterName, _ := reader.ReadString('\n')

	fmt.Print("Name an actual location: ")
	location, _ := reader.ReadString('\n')

	fmt.Print("What's your favorite food? ")
	food, _ := reader.ReadString('\n')

	// Compose and display the story
	fmt.Printf("Recently, %s was seen in %s eating %s!\n", characterName, location, food)
}

