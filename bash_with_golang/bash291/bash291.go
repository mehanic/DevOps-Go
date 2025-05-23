package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Function to convert string to lowercase
func toLower(input string) string {
	return strings.ToLower(input)
}

// Function to perform backup
func doBackup() {
	homeDir, _ := os.UserHomeDir()
	cmd := exec.Command("tar", "-czvf", homeDir+"/backup.tgz", homeDir+"/bin")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Backup failed:", err)
	}
}

// Function to display calendar
func showCal() {
	var cmd *exec.Cmd

	// Check if ncal exists, otherwise fallback to cal
	if _, err := os.Stat("/usr/bin/ncal"); err == nil {
		cmd = exec.Command("/usr/bin/ncal", "-w")
	} else {
		cmd = exec.Command("/usr/bin/cal")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Calendar display failed:", err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")

		// Menu display
		fmt.Println("Choose an item: a, b or c")
		fmt.Println("a: Backup")
		fmt.Println("b: Display Calendar")
		fmt.Println("c: Exit")
		fmt.Print("Your choice: ")

		// Read single character input
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Convert to lowercase
		input = toLower(input)

		// Handle the input choices
		switch input {
		case "a":
			doBackup()
		case "b":
			showCal()
		case "c":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}

		// Prompt to continue
		fmt.Print("Press any key to continue...")
		reader.ReadString('\n')
	}
}

