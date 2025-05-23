package main

import (
//	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// 1. Simulate grep --color=auto functionality (highlight match)
	fmt.Println("Grep with color highlighting:")
	highlighted := grepWithHighlight("this is a word\nnext line", "word")
	fmt.Println(highlighted)

	// 2. Simulate grep -c (count matching lines)
	fmt.Println("\nGrep -c (count matching lines):")
	count := grepCount("[0-9]", "1 2 3 4\nhello\n5 6")
	fmt.Println(count)

	// 3. Simulate grep -n (show matching line numbers)
	fmt.Println("\nGrep -n (show matching line numbers):")
	matches := grepWithLineNumbers("[0-9]", "1 2 3 4\nhello\n5 6")
	fmt.Println(matches)

	// 4. Simulate grep -b (show byte offsets of matches)
	fmt.Println("\nGrep -b (show byte offsets):")
	byteOffsets := grepByteOffsets("not", "gnu is not unix")
	fmt.Println(byteOffsets)

	// 5. Simulate grep -l (list matching files)
	fmt.Println("\nGrep -l (list matching files):")
	createTestFiles()
	matchingFiles := grepMatchingFiles("hello", []string{"1.txt", "2.txt", "3.txt"})
	fmt.Println(strings.Join(matchingFiles, "\n"))
}

// Function to simulate grep with color highlighting
func grepWithHighlight(text, pattern string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		return fmt.Sprintf("\033[31m%s\033[0m", match) // Highlight in red
	})
}

// Function to simulate grep -c (count matching lines)
func grepCount(pattern, text string) int {
	re := regexp.MustCompile(pattern)
	lines := strings.Split(text, "\n")
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

// Function to simulate grep -n (show matching line numbers)
func grepWithLineNumbers(pattern, text string) string {
	re := regexp.MustCompile(pattern)
	lines := strings.Split(text, "\n")
	var result strings.Builder
	for i, line := range lines {
		if re.MatchString(line) {
			result.WriteString(fmt.Sprintf("%d:%s\n", i+1, line))
		}
	}
	return result.String()
}

// Function to simulate grep -b (show byte offsets of matches)
func grepByteOffsets(pattern, text string) string {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringIndex(text, -1)
	var result strings.Builder
	for _, match := range matches {
		offset := match[0]
		result.WriteString(fmt.Sprintf("%d:%s\n", offset, text[match[0]:match[1]]))
	}
	return result.String()
}

// Function to simulate grep -l (list matching files)
func grepMatchingFiles(pattern string, files []string) []string {
	re := regexp.MustCompile(pattern)
	var matchingFiles []string

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}
		if re.Match(content) {
			matchingFiles = append(matchingFiles, file)
		}
	}
	return matchingFiles
}

// Helper function to create test files
func createTestFiles() {
	file1 := []byte("hello\nliyang\n")
	file2 := []byte("hello\nliuyuanyuan\n")
	file3 := []byte("goodbye\nliyang\n")

	os.WriteFile("1.txt", file1, 0644)
	os.WriteFile("2.txt", file2, 0644)
	os.WriteFile("3.txt", file3, 0644)
}

