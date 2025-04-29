package main

import (
	"bufio"
	"fmt"
	"os"
)

func rwText() {
	// Read the file line by line
	file, err := os.Open("somefile.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// process line
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Write chunks of text data
	file, err = os.Create("somefile.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("text1")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	_, err = file.WriteString("text2")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func main() {
	rwText()
}

