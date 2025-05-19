package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
//	"strings"
)

func main() {
	// Check the number of arguments
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s filename string_length\n", os.Args[0])
		os.Exit(1)
	}

	// Read filename and string length
	filename := os.Args[1]
	strLength, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: string_length must be an integer")
		os.Exit(1)
	}

	// Calculate half the string length
	count := strLength / 2

	// Build the base pattern for the regex
	basePattern := "/^"
	for i := 0; i < count; i++ {
		basePattern += `\(.\)`
	}

	// If string length is odd, add an extra character
	if strLength%2 != 0 {
		basePattern += "."
	}

	// Reverse part
	for count > 0 {
		basePattern += `\` + strconv.Itoa(count)
		count--
	}
	basePattern += "$/p"

	// Output the generated base pattern
	fmt.Println(basePattern)

	// Execute the sed command
	cmd := exec.Command("sed", "-n", basePattern, filename)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing sed:", err)
		os.Exit(1)
	}

	// Print the output from sed
	fmt.Println(string(output))
}

