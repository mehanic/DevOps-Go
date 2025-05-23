package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Ask for the user's name
		fmt.Println("Who are you?")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name) // Remove newline characters

		if name != "Joe" {
			continue
		}

		// Ask for the password
		fmt.Println("Hello, Joe. What is the password? (It is a fish.)")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password) // Remove newline characters

		if password == "swordfish" {
			break
		}
	}

	fmt.Println("Access granted.")
}

