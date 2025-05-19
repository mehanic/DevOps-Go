package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Prompt the user for the server name
	var serverName string
	fmt.Print("Which server do you want to connect to: ")
	fmt.Scan(&serverName)

	// Prompt the user for the username
	var userName string
	fmt.Print("Which username do you want to use: ")
	fmt.Scan(&userName)

	// Create the SSH command
	sshCommand := fmt.Sprintf("ssh %s@%s", userName, serverName)

	// Execute the SSH command
	cmd := exec.Command("bash", "-c", sshCommand)

	// Run the command and capture the output
	cmd.Stdout = exec.Command("tee", "/dev/tty").Stdout
	cmd.Stderr = exec.Command("tee", "/dev/tty").Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}

