package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
//	"strconv"
//	"strings"
)

func main() {
	// 1. Generating a sequence of numbers from 1 to 100 and printing lines 50 to 60
	for i := 1; i <= 100; i++ {
		if i >= 50 && i <= 60 {
			fmt.Println(i)
		}
	}

	// 2. Creating a file "1.txt" and writing lines into it
	fileContent := `line with pattern1
line with pattern2
line with pattern3
line end with pattern4
line with pattern5`

	err := ioutil.WriteFile("1.txt", []byte(fileContent), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File '1.txt' created and written successfully.")

	// 3. Reading the file and printing lines between the pattern1 and "end"
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var printLines bool
	startPattern := regexp.MustCompile(`pattern1`)
	endPattern := regexp.MustCompile(`end`)

	for scanner.Scan() {
		line := scanner.Text()

		// Check for start and end patterns
		if startPattern.MatchString(line) {
			printLines = true
		}
		if printLines {
			fmt.Println(line)
		}
		if endPattern.MatchString(line) {
			printLines = false
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

