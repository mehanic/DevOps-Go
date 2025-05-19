package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	choice := 4
	reader := bufio.NewReader(os.Stdin)

	for choice == 4 {
		fmt.Println("1. Bash")
		fmt.Println("2. Scripting")
		fmt.Println("3. Tutorial")
		fmt.Print("Please choose a word [1,2 or 3]? ")

		// Read user input
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Convert input to an integer
		userChoice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// Nested if-else structure
		if userChoice == 1 {
			fmt.Println("You have chosen word: Bash")
			choice = 1
		} else if userChoice == 2 {
			fmt.Println("You have chosen word: Scripting")
			choice = 2
		} else if userChoice == 3 {
			fmt.Println("You have chosen word: Tutorial")
			choice = 3
		} else {
			fmt.Println("Please make a choice between 1-3!")
			choice = 4
		}
	}
}

