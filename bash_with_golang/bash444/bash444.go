package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Display prompt
	fmt.Println("Do you wish to continue? (yes/no)")

	// Read input from user
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	// Check user's answer
	if input == "yes" || input == "y" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

