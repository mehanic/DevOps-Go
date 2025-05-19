package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Check if a filename is provided as an argument
	if len(os.Args) < 2 {
		fmt.Printf("The input to %s should be a filename\n", os.Args[0])
		os.Exit(1)
	}

	// Get the filename from the first argument
	filename := os.Args[1]

	// Check if the provided filename exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", filename)
		os.Exit(1)
	}

	// Open the file for reading server names
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Open or create the output file
	outFile, err := os.Create("server.out")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	// Write the header to the output file
	currentDate := time.Now().Format("01/02/2006")
	fmt.Fprintf(outFile, "The following servers are up on %s\n", currentDate)

	// Scan the input file line by line (each line is a server name)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		server := scanner.Text()

		// Execute the ping command to check if the server is up
		err := exec.Command("ping", "-c1", server).Run()
		if err == nil {
			// If ping succeeds, write the server name to the output file
			fmt.Fprintf(outFile, "Server up: %s\n", server)
		}
	}

	// Check for any errors encountered while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Output the contents of the result file
	outFile.Sync()
	fmt.Println("Ping results:")
	cmd := exec.Command("cat", "server.out")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

