package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	// Create a multi-line string with the fruit names
	fruits := `cherry
mango
apple
banana`

	// Execute the sort command
	cmd := exec.Command("sort")

	// Use a pipe to pass the fruit names to the sort command
	cmd.Stdin = bytes.NewBufferString(fruits)

	// Get the output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing sort command:", err)
		return
	}

	// Print the sorted output
	fmt.Println(string(output))
}

