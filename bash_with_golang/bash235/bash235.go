package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var serverAddr string

	// Prompt for server address
	fmt.Print("Which server should be pinged: ")
	fmt.Scanln(&serverAddr)

	// Execute the ping command
	cmd := exec.Command("ping", "-c", "3", serverAddr)

	// Suppress output by setting Stdout and Stderr to nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	// Run the command and check for errors
	err := cmd.Run()
	if err != nil {
		// If there's an error, print "Server Dead"
		fmt.Println("Server Dead")
	} else {
		fmt.Println("Ping successful.")
	}
}

