package main

import (
	"fmt"
	"io"
	"os"
)

const (
	fileSize   = 100 * 1024 // 100 KB
	fileCount  = 3          // Number of 100 KB files to create
	splitSize  = 100 * 1024 // 100 KB for split files
)

func main() {
	// Step 1: Create a file filled with zeros
	dataFileName := "data.file"
	err := createZeroFilledFile(dataFileName, fileSize, fileCount)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}

	// Step 2: Split the created file into smaller files
	err = splitFile(dataFileName, splitSize)
	if err != nil {
		fmt.Printf("Error splitting file: %v\n", err)
		return
	}

	fmt.Println("File created and split successfully.")
}

// createZeroFilledFile creates a file filled with zeros
func createZeroFilledFile(filename string, size int, count int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write zeros to the file
	zeros := make([]byte, size) // Create a byte slice of the specified size
	for i := 0; i < count; i++ {
		if _, err := file.Write(zeros); err != nil {
			return err
		}
	}

	return nil
}

// splitFile splits a file into smaller files of specified size
func splitFile(filename string, splitSize int) error {
	// Open the source file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var partNumber int
	for {
		// Create a new split file
		splitFileName := fmt.Sprintf("split_data.%04d", partNumber)
		splitFile, err := os.Create(splitFileName)
		if err != nil {
			return err
		}

		// Read data from the original file and write to the split file
		_, err = io.CopyN(splitFile, file, int64(splitSize))
		if err != nil && err != io.EOF {
			splitFile.Close() // Close split file before returning on error
			return err
		}

		splitFile.Close()
		partNumber++

		// Break the loop if EOF is reached
		if err == io.EOF {
			break
		}
	}

	return nil
}

