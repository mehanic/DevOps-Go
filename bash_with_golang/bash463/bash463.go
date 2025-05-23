package main

import (
	"fmt"
	"os"
)

func main() {
	// Ensure that the required arguments are provided
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run story.go <character_name> <location> <food>")
		return
	}

	// Capture command-line arguments
	characterName := os.Args[1]
	location := os.Args[2]
	food := os.Args[3]

	// Compose and display the story
	fmt.Printf("Recently, %s was seen in %s eating %s!\n", characterName, location, food)
}

