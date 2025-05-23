package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a scanner for reading input from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Print options for preferred programming language
	fmt.Println("What is your preferred programming / scripting language")
	fmt.Println("1) bash")
	fmt.Println("2) perl")
	fmt.Println("3) python")
	fmt.Println("4) c++")
	fmt.Println("5) I do not know !")

	// Prompt user for input
	fmt.Print("Please make a choice: ")
	scanner.Scan() // Read input
	choose := scanner.Text() // Get the text input

	// Use a switch case to respond based on user input
	switch choose {
	case "1":
		fmt.Println("You selected bash")
	case "2":
		fmt.Println("You selected perl")
	case "3":
		fmt.Println("You selected python")
	case "4":
		fmt.Println("You selected c++")
	case "5":
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid choice, please try again.")
	}
}

