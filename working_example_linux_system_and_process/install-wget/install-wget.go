package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Command to install wget
	cmd := exec.Command("sudo", "apt-get", "install", "-y", "wget")

	// Run the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// Print the output of the command
	fmt.Printf("Output: %s\n", output)
}
