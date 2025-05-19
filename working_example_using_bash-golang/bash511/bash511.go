package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Handle minor errors.
func handleMinorError() {
	fmt.Println("A minor error has occurred, please check the output.")
}

// Handle fatal errors and exit the program.
func handleFatalError() {
	fmt.Println("A critical error has occurred, stopping the script.")
	os.Exit(1)
}

// Function to execute a shell command and handle the result.
func runCommand(cmd string, fatal bool) {
	// Execute the command using "sh -c" to allow shell syntax.
	command := exec.Command("sh", "-c", cmd)
	err := command.Run()

	// Handle errors based on the severity.
	if err != nil {
		if fatal {
			handleFatalError()
		} else {
			handleMinorError()
		}
	}
}

func main() {
	// Minor failures: These will trigger the minor error handler if they fail.
	runCommand("ls -l /tmp/", false)
	runCommand("ls -l /root/", false)

	// Fatal failures: These will stop the program if they fail.
	runCommand("cat /etc/shadow", true)
	runCommand("cat /etc/passwd", true)
}

