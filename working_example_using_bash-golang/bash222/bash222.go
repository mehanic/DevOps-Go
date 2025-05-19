package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Execute the df -h command
	cmd := exec.Command("df", "-h")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Split the output into lines
	lines := strings.Split(out.String(), "\n")

	// Iterate over the lines and print the fourth column
	for _, line := range lines {
		columns := strings.Fields(line) // Split line into fields
		if len(columns) >= 4 { // Ensure there are at least 4 columns
			fmt.Println(columns[3]) // Print the 4th column (available space)
		}
	}
}

