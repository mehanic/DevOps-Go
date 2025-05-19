package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}

	filePath := os.Args[1]

	// Execute the 'ls -l' command
	cmd := exec.Command("ls", "-l", filePath)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing ls: %v\n", err)
		return
	}

	// Split output into lines and fields
	lines := strings.Split(string(output), "\n")
	if len(lines) < 1 {
		fmt.Println("No output from ls command.")
		return
	}

	// The first line contains the file details
	line := lines[0]
	parts := strings.Fields(line)

	if len(parts) < 9 {
		fmt.Println("Unexpected output format from ls command.")
		return
	}

	// Extract required details
	lcount := parts[1] // Link count
	size := parts[4]   // Size
	file := strings.Join(parts[8:], " ") // File name (in case of spaces)

	// Output the results
	fmt.Printf("%s has %s link(s) and is %s bytes long.\n", file, lcount, size)
}

