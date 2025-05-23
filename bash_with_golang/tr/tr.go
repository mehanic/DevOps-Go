package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"
)

// helper function to replace digits with letters
func replaceDigitsWithLetters(input string) string {
	var buffer bytes.Buffer
	for _, r := range input {
		if unicode.IsDigit(r) {
			buffer.WriteRune('a' + (r - '0'))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// helper function to remove characters not in set (complement) and replace digits
func trDeleteComplement(input string) string {
	// Keep only digits and newlines
	re := regexp.MustCompile(`[^0-9\n]+`)
	result := re.ReplaceAllString(input, "")
	return replaceDigitsWithLetters(result)
}

// helper function to compress consecutive spaces
func compressSpaces(input string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(input, " ")
}

// helper function to sum the numbers from the sum.txt file
func sumFileContents(filename string) (int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	content := strings.TrimSpace(string(data))
	numbers := strings.Split(content, "\n")

	sum := 0
	for _, numStr := range numbers {
		var num int
		fmt.Sscanf(numStr, "%d", &num)
		sum += num
	}
	return sum, nil
}

// main function
func main() {
	// Open the output file
	outputFile, err := os.Create("tr.output")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// 1. tr -d -c '0-9\n' | tr '0-9' 'a-z'
	input1 := "hello 123 world\n"
	result1 := trDeleteComplement(input1)
	outputFile.WriteString(result1 + "\n")

	// 2. tr -s ' '
	input2 := "GNU is not UNIX. Recursive right ?"
	result2 := compressSpaces(input2)
	outputFile.WriteString(result2 + "\n")

	// 3. Use tr for addition simulation and append sum
	outputFile.WriteString("sum.txt : \n")
	sum, err := sumFileContents("sum.txt")
	if err != nil {
		fmt.Println("Error reading sum.txt:", err)
		return
	}
	outputFile.WriteString(fmt.Sprintf("%d\n", sum))

	// 4. tr '[:lower:]' '[:upper:]'
	input4 := "hello world"
	result4 := strings.ToUpper(input4)
	outputFile.WriteString(result4 + "\n")

	// Display output file with line numbers
	outputFile.Seek(0, 0)
	scanner := bufio.NewScanner(outputFile)
	lineNumber := 1
	fmt.Println("Output file content:")
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", lineNumber, scanner.Text())
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading output file:", err)
	}
}

