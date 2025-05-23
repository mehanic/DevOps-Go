package main

import (
	"fmt"
	"os"
	"github.com/manifoldco/promptui"
)

func main() {
	// Create a confirmation prompt
	prompt := promptui.Select{
		Label: "Do you wish to continue?",
		Items: []string{"Yes", "No"},
	}

	// Run the prompt and handle the result
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	// Check the user's choice and respond accordingly
	if result == "Yes" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

