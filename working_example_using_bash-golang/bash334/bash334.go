package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	// Create the message to be sent
	message := `Tomorrow, on Friday evening, we will be celebrating
Birthday of few of our colleagues.
All are requested to be present in cafeteria by 3.30 PM.
John`

	// Execute the wall command
	cmd := exec.Command("wall")

	// Use a pipe to pass the message input to the wall command
	cmd.Stdin = bytes.NewBufferString(message)

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	// Print confirmation
	fmt.Println("Message sent")
}

