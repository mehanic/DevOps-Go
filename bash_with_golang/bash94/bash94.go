package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
//	"strings"
)

// Function to create a file with given content
func createFile(filename string, content []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	for _, line := range content {
		file.WriteString(line + "\n")
	}
}

// Function to read a file and return its content as a slice of strings
func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Function to sort and overwrite the file
func sortFile(filename string) error {
	lines, err := readFile(filename)
	if err != nil {
		return err
	}

	sort.Strings(lines)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		file.WriteString(line + "\n")
	}
	return nil
}

// Function to compute intersection of two sorted string slices
func intersection(a, b []string) []string {
	result := []string{}
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			result = append(result, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

// Function to compute union of two sorted string slices
func union(a, b []string) []string {
	result := []string{}
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			result = append(result, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	for i < len(a) {
		result = append(result, a[i])
		i++
	}

	for j < len(b) {
		result = append(result, b[j])
		j++
	}

	return result
}

func main() {
	// Create a.txt and b.txt with given content
	createFile("a.txt", []string{"apple", "orange", "gold", "silver", "steer", "iron"})
	createFile("b.txt", []string{"orange", "gold", "cookies", "carrot"})

	// Sort the files a.txt and b.txt
	sortFile("a.txt")
	sortFile("b.txt")

	// Read the sorted files
	aContent, _ := readFile("a.txt")
	bContent, _ := readFile("b.txt")

	// Find the intersection (comm without -1 -2)
	fmt.Println("Intersection:")
	for _, line := range intersection(aContent, bContent) {
		fmt.Println(line)
	}

	fmt.Println("------------------")

	// Find the union (comm with -1 -2)
	fmt.Println("Union:")
	for _, line := range union(aContent, bContent) {
		fmt.Println(line)
	}
}

