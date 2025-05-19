package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
//	"strconv"
)

// Function to print lines 50 to 60 from 1 to 100
func printSeq() {
	for i := 1; i <= 100; i++ {
		if i >= 50 && i <= 60 {
			fmt.Println(i)
		}
	}
}

// Function to create a file and write content to it
func createFile() {
	content := `line with pattern1
line with pattern2
line with pattern3
line end with pattern4
line with pattern5
`
	err := os.WriteFile("1.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
}

// Function to simulate the awk search '/pattern1/, /end/' from file
func searchPattern() {
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inRange := false
	startPattern := regexp.MustCompile("pattern1")
	endPattern := regexp.MustCompile("end")

	for scanner.Scan() {
		line := scanner.Text()
		if startPattern.MatchString(line) {
			inRange = true
		}
		if inRange {
			fmt.Println(line)
		}
		if endPattern.MatchString(line) {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func main() {
	// Equivalent to 'seq 1 100 | awk 'NR==50, NR==60''
	fmt.Println("Printing sequence from 50 to 60:")
	printSeq()

	// Equivalent to 'touch 1.txt' and 'cat > 1.txt <<EOF'
	createFile()

	// Equivalent to 'awk '/pattern1/, /end/' 1.txt'
	fmt.Println("\nLines between pattern1 and end:")
	searchPattern()
}

