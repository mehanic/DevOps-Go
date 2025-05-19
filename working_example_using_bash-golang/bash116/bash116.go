package main

import (
	"fmt"
	"io/ioutil"
	"os"
//	"os/exec"
	"sort"
	"strings"
)

func main() {
	// Step 1: Remove data.txt if it exists
	if _, err := os.Stat("data.txt"); err == nil {
		err := os.Remove("data.txt")
		if err != nil {
			fmt.Println("Error removing data.txt:", err)
			return
		}
	}

	// Step 2: Create a new data.txt file and write to it
	data := []string{
		"1 max 2000",
		"2 winxp 4000",
		"3 bsd 1000",
		"4 linux 1000",
	}

	err := ioutil.WriteFile("data.txt", []byte(strings.Join(data, "\n")), 0644)
	if err != nil {
		fmt.Println("Error writing to data.txt:", err)
		return
	}

	// Step 3: Sort the contents and print to stdout
	fmt.Println("Sorted contents:")
	content, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading data.txt:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	sort.Strings(lines)

	for _, line := range lines {
		fmt.Println(line)
	}

	// Step 4: Sort and save the sorted output to a file named log
	err = ioutil.WriteFile("log", []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		fmt.Println("Error writing to log:", err)
		return
	}

	// Step 5: Sort numerically based on the first column
	numericSort(lines)
	fmt.Println("\nNumerically sorted by the first column:")
	for _, line := range lines {
		fmt.Println(line)
	}

	// Step 6: Sort based on the second column
	sort.Slice(lines, func(i, j int) bool {
		partsI := strings.Fields(lines[i])
		partsJ := strings.Fields(lines[j])
		return partsI[1] < partsJ[1]
	})
	fmt.Println("\nSorted by the second column:")
	for _, line := range lines {
		fmt.Println(line)
	}

	// Step 7: Sort numerically based on the first column and print
	sort.SliceStable(lines, func(i, j int) bool {
		partsI := strings.Fields(lines[i])
		partsJ := strings.Fields(lines[j])
		return partsI[0] < partsJ[0]
	})
	fmt.Println("\nNumerically sorted by the first column (stable):")
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Function to perform numeric sort on lines based on the first field
func numericSort(lines []string) {
	sort.Slice(lines, func(i, j int) bool {
		partsI := strings.Fields(lines[i])
		partsJ := strings.Fields(lines[j])
		return partsI[0] < partsJ[0]
	})
}

