package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Define color constants
const (
	GREEN  = "\033[32m"
	RESET  = "\033[0m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for color or mono
	fmt.Print("Type color or mono for script output: ")
	reply, _ := reader.ReadString('\n')
	reply = strings.TrimSpace(reply) // Remove newline character

	// Check if the reply matches "color" or "colour"
	matched, err := regexp.MatchString(`(?i)colou?r`, reply) // case insensitive match
	if err != nil {
		fmt.Println("Error in regex:", err)
		return
	}

	// If it matches, we can assume we need to load color configurations
	if matched {
		if err := loadColorSettings(); err != nil {
			fmt.Println("Error loading color settings:", err)
			return
		}
	}

	// Print message in green if color is selected
	fmt.Printf("%sThis is %s%s\n", GREEN, os.Args[0], RESET)

	// Exit the program
	os.Exit(0)
}

// loadColorSettings simulates sourcing a configuration file
func loadColorSettings() error {
	// In a real implementation, you might read from a file or set environment variables
	// Here, we will just simulate this action
	fmt.Println("Color settings loaded.")
	return nil
}

