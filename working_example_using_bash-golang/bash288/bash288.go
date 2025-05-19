package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// Check if a filename argument is provided
	if len(os.Args) < 2 {
		fmt.Printf("The input to %s should be a filename\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("The file %s does not exist.\n", filename)
		os.Exit(1)
	}

	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Open the output file
	outputFile, err := os.Create("server.out")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// Write the header to the output file
	currentDate := time.Now().Format("01/02/2006")
	outputFile.WriteString(fmt.Sprintf("The following servers are up on %s\n", currentDate))

	// Create a buffered reader for the input file
	scanner := bufio.NewScanner(file)

	// Loop through each line (server address)
	for scanner.Scan() {
		server := strings.TrimSpace(scanner.Text())
		if server == "" {
			continue
		}

		// Ping the server
		cmd := exec.Command("ping", "-c1", server)
		err := cmd.Run()

		// If the server is up, write to the output file
		if err == nil {
			outputFile.WriteString(fmt.Sprintf("Server up: %s\n", server))
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	// Display the contents of the output file
	content, err := os.ReadFile("server.out")
	if err != nil {
		fmt.Println("Error reading the output file:", err)
		os.Exit(1)
	}
	fmt.Println(string(content))
}

