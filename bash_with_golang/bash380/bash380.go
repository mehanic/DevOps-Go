package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Equivalent of whoami (returns the current username)
	currentUser, err := exec.Command("whoami").Output()
	if err != nil {
		fmt.Println("Error executing whoami:", err)
	} else {
		fmt.Printf("Current user: %s", currentUser)
	}

	// Equivalent of pwd (returns the current working directory)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	} else {
		fmt.Printf("Current working directory: %s\n", currentDir)
	}

	// Equivalent of ls (lists the contents of the current directory)
	lsCmd := exec.Command("ls")
	lsOutput, err := lsCmd.Output()
	if err != nil {
		fmt.Println("Error executing ls:", err)
	} else {
		fmt.Printf("Directory listing:\n%s", lsOutput)
	}

	// Equivalent of echo commands (prints messages)
	fmt.Println("Echo one 1")
	fmt.Println("Echo two 2")
}

