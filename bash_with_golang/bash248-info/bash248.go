package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// toLower converts the input string to lowercase.
func toLower(input string) string {
	return strings.ToLower(input)
}

// doBackup performs a backup by creating a tar.gz archive of the ~/bin directory.
func doBackup() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}
	backupFile := homeDir + "/backup.tgz"
	cmd := exec.Command("tar", "-czvf", backupFile, homeDir+"/bin")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error during backup:", err)
	}
}

// showCal displays the calendar using either `ncal -w` or `cal`.
func showCal() {
	if _, err := os.Stat("/usr/bin/ncal"); err == nil {
		// Use ncal if available
		cmd := exec.Command("/usr/bin/ncal", "-w")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else {
		// Fallback to cal
		cmd := exec.Command("/usr/bin/cal")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Clear the terminal screen
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		// Display the menu
		fmt.Println("Choose an item: a, b, or c")
		fmt.Println("a: Backup")
		fmt.Println("b: Display Calendar")
		fmt.Println("c: Exit")

		// Read single character input
		char, _ := reader.ReadString('\n')
		char = strings.TrimSpace(char)
		char = toLower(char)

		// Handle the selected option
		switch char {
		case "a":
			doBackup()
		case "b":
			showCal()
		case "c":
			os.Exit(0)
		}

		// Wait for the user to press any key before continuing
		fmt.Println("Press Enter to continue...")
		reader.ReadString('\n')
	}
}

