package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Display menu
		fmt.Println("1. Show disk usage.")
		fmt.Println("2. Show system uptime.")
		fmt.Println("3. Show the users logged into the system.")
		fmt.Print("What would you like to do? (Enter q to quit): ")

		// Read user input
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// Handle user choice
		switch choice {
		case "1":
			// Run 'df' to show disk usage
			runCommand("df")
		case "2":
			// Run 'uptime' to show system uptime
			runCommand("uptime")
		case "3":
			// Run 'who' to show logged-in users
			runCommand("who")
		case "q":
			// Quit the loop
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option.")
		}
		fmt.Println()
	}
}

// runCommand executes a system command and prints the output
func runCommand(command string) {
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running command %s: %v\n", command, err)
	}
}

