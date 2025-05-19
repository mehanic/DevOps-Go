package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	// Create a multi-line string with the text input
	text := `There was major earthquake
On April 25, 2015
in Nepal.
There was huge loss of human life in this tragic event.`

	// Execute the wc command with the -w flag
	cmd := exec.Command("wc", "-w")

	// Use a pipe to pass the text input to the wc command
	cmd.Stdin = bytes.NewBufferString(text)

	// Get the output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing wc command:", err)
		return
	}

	// Print the word count
	fmt.Printf("Word count: %s", output)
}

