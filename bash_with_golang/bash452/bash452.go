package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput() // Combine stdout and stderr
	if err != nil {
		return fmt.Errorf("error executing command '%s %s': %v\nOutput: %s", command, args, err, output)
	}
	fmt.Print(string(output)) // Print the output to stdout
	return nil
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		// Windows
		execCommand("cmd", "/c", "cls")
	} else {
		// Unix-based systems (Linux, MacOS)
		execCommand("clear")
	}
}

func main() {
	// Clear the screen
	clearScreen()

	// List files in long format
	fmt.Println("Listing files in long format:")
	if err := execCommand("ls", "-l"); err != nil {
		fmt.Println(err)
		return
	}

	// Display the current date and time
	fmt.Println("Current date and time:")
	if err := execCommand("date"); err != nil {
		fmt.Println(err)
		return
	}

	// Install python3
	fmt.Println("Installing python3...")
	if err := execCommand("sudo", "apt", "install", "-y", "python3"); err != nil {
		fmt.Println(err)
		return
	}
}

