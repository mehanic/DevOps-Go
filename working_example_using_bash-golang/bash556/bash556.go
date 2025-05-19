package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Open output file
	outputFile, err := os.Create("exampleXargs.output")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// 1. Split and process the string "solitXsplitXsplitXsplit"
	firstInput := "solitXsplitXsplitXsplit"
	parts := strings.Split(firstInput, "X")

	// Process 2 parts at a time
	for i := 0; i < len(parts); i += 2 {
		group := strings.Join(parts[i:min(i+2, len(parts))], " ")
		outputFile.WriteString(group + "\n")
	}
	outputFile.WriteString("\n")

	// 2. Read and process content from exampleXargs.input with default grouping
	inputContent, err := ioutil.ReadFile("exampleXargs.input")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Process the input with default grouping (all at once)
	args := strings.Fields(string(inputContent))
	outputFile.WriteString(strings.Join(args, " ") + "\n")
	outputFile.WriteString("\n")

	// 3. Read and process content from exampleXargs.input with grouping of 3 arguments at a time
	for i := 0; i < len(args); i += 3 {
		group := strings.Join(args[i:min(i+3, len(args))], " ")
		outputFile.WriteString(group + "\n")
	}
	outputFile.WriteString("\n")

	// 4. Display the output with line numbers
	output, err := ioutil.ReadFile("exampleXargs.output")
	if err != nil {
		fmt.Println("Error reading output file:", err)
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	lineNumber := 1
	for scanner.Scan() {
		fmt.Printf("%d\t%s\n", lineNumber, scanner.Text())
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}

// min helper function to handle bounds
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

