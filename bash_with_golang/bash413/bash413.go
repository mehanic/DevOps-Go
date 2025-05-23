package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var iname string
var oname string
var fname string
var where int

func determineTypeOfFile(file string) bool {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer f.Close()

	buf := make([]byte, 512)
	_, err = f.Read(buf)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return false
	}

	// Check if the file is ASCII
	isASCII := true
	for _, b := range buf {
		if b > 127 {
			isASCII = false
			break
		}
	}

	if isASCII {
		fmt.Println("ASCII file - continuing")
	} else {
		fmt.Println("Not an ASCII file, perhaps it is Binary?")
	}
	return isASCII
}

func openInputAndInsert(input string, block string, finalOutput string, ptr int) error {
	inFile, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("error opening input file: %w", err)
	}
	defer inFile.Close()

	blockFile, err := os.Open(block)
	if err != nil {
		return fmt.Errorf("error opening block file: %w", err)
	}
	defer blockFile.Close()

	// Create a temporary file to store the result
	tmpFile, err := os.CreateTemp("", "temp_output")
	if err != nil {
		return fmt.Errorf("error creating temp file: %w", err)
	}
	defer tmpFile.Close()

	scanner := bufio.NewScanner(inFile)
	var ctr int
	var buffer strings.Builder

	// Read the input file line by line and insert block at the specified line
	for scanner.Scan() {
		line := scanner.Text()
		if ctr == ptr {
			// Append block content at the specified position
			blockScanner := bufio.NewScanner(blockFile)
			for blockScanner.Scan() {
				buffer.WriteString(blockScanner.Text() + "\n")
			}
		}
		buffer.WriteString(line + "\n")
		ctr++
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input file: %w", err)
	}

	// Write the buffer to the temp file
	if _, err := tmpFile.WriteString(buffer.String()); err != nil {
		return fmt.Errorf("error writing to temp file: %w", err)
	}

	// Move the temp file to the final output
	if err := os.Rename(tmpFile.Name(), finalOutput); err != nil {
		return fmt.Errorf("error moving temp file to final output: %w", err)
	}

	return nil
}

func main() {
	// Parse command line arguments
	flag.StringVar(&iname, "i", "", "Input file name (required)")
	flag.StringVar(&oname, "o", "", "Block file name (required)")
	flag.StringVar(&fname, "f", "", "Final output file name (required)")
	flag.IntVar(&where, "w", 0, "Line number where the block should be inserted (default: append)")

	flag.Parse()

	// Validate input parameters
	if iname == "" {
		fmt.Println("ERROR: -i input file parameter required")
		os.Exit(1)
	}

	if oname == "" {
		fmt.Println("ERROR: -o block file parameter required")
		os.Exit(1)
	}

	if fname == "" {
		fmt.Println("ERROR: -f final output file parameter required")
		os.Exit(1)
	}

	if where < 0 {
		fmt.Println("ERROR: -w must be a non-negative number")
		os.Exit(1)
	}

	// Ensure the input file is ASCII text
	if !determineTypeOfFile(iname) {
		os.Exit(1)
	}

	// Insert the block content into the input file and write to the final output
	err := openInputAndInsert(iname, oname, fname, where)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Println("File processed successfully!")
}

