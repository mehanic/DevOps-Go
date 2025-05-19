package main

import (
	"fmt"
	"strings"
)

func main() {
	// The input line
	line := "root:x:0:0:root:/root:/bin/bash"

	// Split the line by colon
	parts := strings.Split(line, ":")

	// Initialize variables to hold the user and shell
	var user, shell string

	// Iterate through the parts of the split string
	for count, item := range parts {
		if count == 0 {
			user = item // First part is the user
		}
		if count == 6 {
			shell = item // Seventh part is the shell
		}
	}

	// Output the result
	fmt.Printf("%s's shell is %s\n", user, shell)
}
