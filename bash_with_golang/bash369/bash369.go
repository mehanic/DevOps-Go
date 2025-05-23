package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Check if a filename was provided as a command-line argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide a filename as a command-line argument.")
	}

	filename := os.Args[1]

	// Read the file content
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		// Print the error to STDERR and exit
		log.Fatalf("Error reading file: %v", err)
	}

	// Print the content of the file to STDOUT
	fmt.Println(string(data))
}

