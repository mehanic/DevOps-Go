package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if the user provided a search term as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <search-term>")
		os.Exit(1)
	}
	searchTerm := os.Args[1]

	// The equivalent of the "here document" using a multi-line string
	data := `
mike x.123
joe x.234
sue x.555
pete x.818
sara x.822
bill x.919
`

	// Use a scanner to read through the string line by line
	scanner := bufio.NewScanner(strings.NewReader(data))

	// Loop through each line and print lines that contain the search term
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchTerm) {
			fmt.Println(line)
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}


//go run here-dokument.go x.919                                      1 (0.180s) < 19:56:07
//bill x.919

