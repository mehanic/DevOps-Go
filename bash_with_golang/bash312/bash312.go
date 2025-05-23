package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Print a message
	fmt.Println("This is a shell script")

	// Run the `ls -lah` command
	cmd := exec.Command("ls", "-lah")
	cmd.Stdout = os.Stdout // Redirect output to standard output
	cmd.Stderr = os.Stderr // Redirect error to standard error
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing ls command:", err)
	}

	// Print another message
	fmt.Println("I am done running ls")

	// Declare and print a variable
	SOMEVAR := "text stuff"
	fmt.Println(SOMEVAR)
}

