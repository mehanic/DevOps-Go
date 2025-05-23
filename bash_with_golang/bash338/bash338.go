package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the file for reading
	file, err := os.Open("sample_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when done

	// Use io.Copy to copy the contents of the file to standard output (like cat <&3 in Bash)
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// The file will automatically be closed when the function exits due to defer
}

