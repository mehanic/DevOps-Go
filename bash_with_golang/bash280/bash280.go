package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Prompt the user for the server address
	var serverAddr string
	fmt.Print("Which server should be pinged: ")
	fmt.Scan(&serverAddr)

	// Execute the ping command
	cmd := exec.Command("ping", "-c", "3", serverAddr)

	// Combine stdout and stderr
	output, err := cmd.CombinedOutput()

	// Check if there was an error
	if err != nil {
		fmt.Println("Server dead")
		return
	}

	// Print the output of the ping command (optional)
	fmt.Println(string(output))
}

