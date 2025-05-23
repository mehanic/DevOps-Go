package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Step 1: Check if the log file exists and remove it
	if _, err := os.Stat("./log"); err == nil {
		err = os.Remove("./log")
		if err != nil {
			fmt.Println("Error deleting file:", err)
		}
	}

	// Step 2: Create and write to the log file
	file, err := os.Create("log")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("1 2 3\n9 8 7 6 3 6\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Step 3: Simulate the behavior of `xargs`
	fmt.Println("Original content in one line (like `xargs`):")
	content, err := os.ReadFile("log")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(strings.ReplaceAll(string(content), "\n", " "))

	fmt.Println("\nSimulate `xargs -n 3` behavior:")
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for i := 0; i < len(fields); i += 3 {
			end := i + 3
			if end > len(fields) {
				end = len(fields)
			}
			fmt.Println(strings.Join(fields[i:end], " "))
		}
	}

	fmt.Println("\nSimulate `xargs -n 1 | xargs -I {} seq 1 {}` behavior:")
	for _, word := range strings.Fields(string(content)) {
		if n, err := strconv.Atoi(word); err == nil {
			for i := 1; i <= n; i++ {
				fmt.Printf("%d ", i)
			}
			fmt.Println()
		}
	}

	// Step 4: Simulate `find` and `xargs` to remove files
	// Creating sample files for the example
	touchFiles := []string{"1.log", "2.log", "3.log"}
	for _, f := range touchFiles {
		err = os.WriteFile(f, []byte("sample"), 0644)
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
	}

	// Using `filepath.Walk` to find .log files and remove them
	fmt.Println("\nDeleting .log files...")
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".log") {
			fmt.Println("Deleting:", path)
			err := os.Remove(path)
			if err != nil {
				fmt.Println("Error deleting file:", err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the file path:", err)
	}
}

