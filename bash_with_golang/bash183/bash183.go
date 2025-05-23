package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Execute the "cat /etc/shadow" command
	cmd := exec.Command("cat", "/etc/shadow")
	err := cmd.Run()

	// Check if the command was successful
	if err == nil {
		fmt.Println("Command succeeded")
		os.Exit(0)
	} else {
		fmt.Println("Command failed")
		os.Exit(1)
	}
}

