package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <FILE_NAME>")
		return
	}

	fileName := os.Args[1]

	// Step 1: Strip underscores
	fileNameClean := strings.ReplaceAll(fileName, "_", "")

	// Step 2: Remove every second character (similar to sed 's/..//g')
	// We can achieve this by iterating through the string
	var resultBuilder strings.Builder
	for i, r := range fileNameClean {
		// Append only characters at even indices (0-based)
		if i%2 == 0 {
			resultBuilder.WriteRune(r)
		}
	}
	fileNameClean = resultBuilder.String()

	// Step 3: Replace spaces with underscores
	fileNameClean = strings.ReplaceAll(fileNameClean, " ", "_")

	// Step 4: Clean out anything that's not alphanumeric or an underscore
	re := regexp.MustCompile(`[^a-zA-Z0-9_.]`)
	fileNameClean = re.ReplaceAllString(fileNameClean, "")

	// Step 5: Check if the file exists
	if _, err := os.Stat(fileNameClean); os.IsNotExist(err) {
		fmt.Printf("File does not exist: %s\n", fileNameClean)
	} else {
		fmt.Printf("File exists: %s\n", fileNameClean)
	}
}

