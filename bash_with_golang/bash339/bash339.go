package main

import (
	"fmt"
	"os"
)

func main() {
	// File name
	fileName := "sample_out.txt"

	// Open the file for both reading and writing (create if doesn't exist, truncate)
	// os.O_RDWR: Open the file for reading and writing.
	// os.O_CREATE: Create the file if it does not exist.
	// os.O_TRUNC: Truncate the file to zero size when opening.
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure file is closed after operations

	// Content to write to the file
	content := `
Do not dwell in the past,
do not dream of the future,
concentrate the mind on the present moment. - Buddha
`

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Rewind to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking to start of file:", err)
		return
	}

	// Read the content back from the file
	buf := make([]byte, 1024) // Create a buffer to read into
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	// Print the content read from the file
	fmt.Println("File content:")
	fmt.Println(string(buf[:n]))
}

