package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define color escape codes as constants
const (
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	BLUE   = "\033[34m"
	RESET  = "\033[0m"
)

func main() {
	// Retrieve the first argument or default to "Anonymous"
	name := "Anonymous"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// Greeting message
	fmt.Printf("Hello %s\n", name)

	// Input prompt for color or mono
	fmt.Print("Type color or mono for script output: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Convert input to lowercase for case-insensitive matching
	input = strings.ToLower(input)

	// By default, disable color
	colorEnabled := false

	// Check if input is color (allow variations like color/colour)
	if strings.Contains(input, "color") || strings.Contains(input, "colour") {
		colorEnabled = true
	}

	// Output based on the user's choice
	if colorEnabled {
		fmt.Printf("%sThis is %s%s\n", GREEN, os.Args[0], RESET)
	} else {
		fmt.Printf("This is %s\n", os.Args[0])
	}

	// Exit the program successfully
	os.Exit(0)
}

