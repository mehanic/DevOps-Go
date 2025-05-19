package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the /etc/passwd file
	file, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Println("Error opening /etc/passwd:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read and print the first line (equivalent to `read line < /etc/passwd` in bash)
	if scanner.Scan() {
		firstLine := scanner.Text()
		fmt.Println(firstLine)
	}

	// Read and print the rest of the lines (equivalent to the `while read` loop in bash)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// Handle any error that occurred during the scan
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading /etc/passwd:", err)
	}
}

