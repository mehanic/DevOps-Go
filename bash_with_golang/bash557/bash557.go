package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Open input file args.input
	inputFile, err := os.Open("args.input")
	if err != nil {
		fmt.Println("Error opening args.input:", err)
		return
	}
	defer inputFile.Close()

	// Open output file cecho.output
	outputFile, err := os.Create("cecho.output")
	if err != nil {
		fmt.Println("Error creating cecho.output:", err)
		return
	}
	defer outputFile.Close()

	// Create a buffered writer for output file
	writer := bufio.NewWriter(outputFile)

	// First part: Execute ./cecho.sh with the arguments from args.input
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		args := strings.Fields(scanner.Text()) // Split input line into arguments

		// Run the ./cecho.sh script with the args
		cmd := exec.Command("./cecho.sh", args...)
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing ./cecho.sh:", err)
			return
		}

		// Write the output to cecho.output
		writer.Write(output)
		writer.Flush()
	}

	// Rewind input file to process again for second part
	inputFile.Seek(0, 0)

	// Second part: Execute ./cecho.sh -p {} -1 for each argument
	for scanner.Scan() {
		arg := scanner.Text() // Read one line at a time

		// Run ./cecho.sh -p {} -1 (replace {} with the current argument)
		cmd := exec.Command("./cecho.sh", "-p", arg, "-1")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing ./cecho.sh with -p", arg, "-1:", err)
			return
		}

		// Append the output to cecho.output
		writer.Write(output)
		writer.Flush()
	}

	// Read and print the output with line numbers (simulating cat -n cecho.output)
	outputFile.Seek(0, 0) // Reset the pointer to the beginning of the file for reading

	scanner = bufio.NewScanner(outputFile)
	lineNumber := 1
	for scanner.Scan() {
		fmt.Printf("%d\t%s\n", lineNumber, scanner.Text())
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading cecho.output:", err)
	}
}

