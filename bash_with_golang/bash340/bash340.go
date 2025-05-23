package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open in_file.txt for reading
	inFile, err := os.Open("in_file.txt")
	if err != nil {
		fmt.Println("Error opening in_file.txt:", err)
		return
	}
	defer inFile.Close() // Ensure the file is closed when done

	// Open out_file.txt for writing (it will be created or truncated)
	outFile, err := os.Create("out_file.txt")
	if err != nil {
		fmt.Println("Error opening out_file.txt:", err)
		return
	}
	defer outFile.Close() // Ensure the file is closed when done

	// Read the first line from in_file.txt
	scanner := bufio.NewScanner(inFile)
	if scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line) // Print the line

		// Write the line to out_file.txt
		_, err := outFile.WriteString(fmt.Sprintf("Line 1 - %s\n", line))
		if err != nil {
			fmt.Println("Error writing to out_file.txt:", err)
			return
		}

		fmt.Println("Writing content of in_file.txt to out_file.txt")
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from in_file.txt:", err)
		return
	}
}

