package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"
)

// Global variables for input arguments
var fname string
var length int
var splitType string

// Function to check if the file is ASCII text
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

	// Check for ASCII content
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

// Function to write buffer to a new file
func writeToFile(baseName string, fileCount int, buffer string) {
	outputFileName := fmt.Sprintf("%s.%d", baseName, fileCount)
	err := os.WriteFile(outputFileName, []byte(buffer), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Printf("Wrote buffer to file: %s\n", outputFileName)
	}
}

// Function to split file by lines
func readAndSplitByLine(baseName string, length int) {
	file, err := os.Open(baseName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	fileCount := 0
	var buffer strings.Builder

	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
		lineCount++

		if lineCount >= length {
			fileCount++
			writeToFile(baseName, fileCount, buffer.String())
			buffer.Reset()
			lineCount = 0
		}
	}

	if buffer.Len() > 0 {
		fileCount++
		writeToFile(baseName, fileCount, buffer.String())
	}
}

// Function to split file by size
func readAndWriteBySize(baseName string, length int) {
	file, err := os.Open(baseName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	pageSize := syscall.Getpagesize()
	buf := make([]byte, pageSize)
	fileCount := 0
	sizeCount := 0
	var buffer strings.Builder

	for {
		n, err := file.Read(buf)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file:", err)
			break
		}

		buffer.Write(buf[:n])
		sizeCount++

		if sizeCount >= length {
			fileCount++
			writeToFile(baseName, fileCount, buffer.String())
			buffer.Reset()
			sizeCount = 0
		}

		if n == 0 {
			break
		}
	}

	if buffer.Len() > 0 {
		fileCount++
		writeToFile(baseName, fileCount, buffer.String())
	}
}

func main() {
	// Parse command line arguments
	flag.StringVar(&fname, "i", "", "Input file name (required)")
	flag.IntVar(&length, "l", 10, "Number of lines or size blocks to split by (default: 10)")
	flag.StringVar(&splitType, "t", "line", "Split type: 'line' or 'size' (default: line)")
	flag.Parse()

	if fname == "" {
		fmt.Println("ERROR: -i input parameter required; filename")
		os.Exit(1)
	}

	if length <= 0 {
		fmt.Println("ERROR: -l must be greater than 0")
		os.Exit(1)
	}

	if splitType != "line" && splitType != "size" {
		fmt.Println("ERROR: -t must be set to 'line' or 'size'")
		os.Exit(1)
	}

	// Make sure the file is ASCII text
	if !determineTypeOfFile(fname) {
		os.Exit(1)
	}

	// Begin splitting the file
	if splitType == "size" {
		readAndWriteBySize(fname, length)
	} else {
		readAndSplitByLine(fname, length)
	}
}

