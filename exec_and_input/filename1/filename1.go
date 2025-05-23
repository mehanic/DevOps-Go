package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <filename>")
		return
	}

	script := os.Args[0]
	filename := os.Args[1]

	// Read the file specified in the command-line argument
	readFile(filename)

	// Read a filename from user input
	fmt.Print("Enter a filename: ")
	var userFilename string
	fmt.Scanln(&userFilename)
	readFile(userFilename)

	// Read a hardcoded filename
	readFile("copied.txt")

	// Read another hardcoded filename
	filename = "ex1.py"
	readFile(filename)
}
