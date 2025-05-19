package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var array []string
	scanner := bufio.NewScanner(file)

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		array = append(array, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the number of elements and their content
	fmt.Printf("Number of elements: %d\n", len(array))
	fmt.Println("Array's content:", array)
}

