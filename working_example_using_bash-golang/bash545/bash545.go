package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Execute the uname -a command using backticks (``)
	output, err := exec.Command("uname", "-a").Output()
	if err != nil {
		fmt.Printf("Error executing command: %s\n", err)
		return
	}
	fmt.Println(string(output))

	// Execute the uname -a command using parentheses (())
	output, err = exec.Command("uname", "-a").Output()
	if err != nil {
		fmt.Printf("Error executing command: %s\n", err)
		return
	}
	fmt.Println(string(output))
}

