package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var serverName string
	var userName string

	// Prompt for server name
	fmt.Print("Which server do you want to connect to: ")
	fmt.Scanln(&serverName)

	// Prompt for username
	fmt.Print("Which username do you want to use: ")
	fmt.Scanln(&userName)

	// Construct the SSH command
	sshCommand := fmt.Sprintf("ssh %s@%s", userName, serverName)

	// Execute the SSH command
	cmd := exec.Command("bash", "-c", sshCommand)
	cmd.Stdout = exec.Command("bash", "-c", "cat").Stdout
	cmd.Stderr = exec.Command("bash", "-c", "cat").Stderr

	// Start the command and check for errors
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
}

