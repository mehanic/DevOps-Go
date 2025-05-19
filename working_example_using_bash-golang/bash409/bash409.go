package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var currentLevel int

// tabCreator creates a string of tabs based on the current level
func tabCreator(level int) string {
	tabs := "."
	for i := 0; i < level; i++ {
		tabs += tabs
	}
	return tabs
}

// recursiveTree walks the file tree starting from the given entry
func recursiveTree(entry string) error {
	return filepath.Walk(entry, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate relative path and current level
		relPath, _ := filepath.Rel(entry, path)
		if relPath == "." {
			return nil // Skip the root directory itself
		}

		// Update current level based on the number of separators in the relative path
		currentLevel = strings.Count(relPath, string(os.PathSeparator))

		// Create tab representation
		tabs := tabCreator(currentLevel)

		if info.IsDir() {
			fmt.Printf("%s_ %s\n", tabs, info.Name())
		} else if info.Mode().IsRegular() {
			fmt.Printf("%s|_%s\n", tabs, info.Name())
		}
		return nil
	})
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}

	parentDir := os.Args[1]

	// Call recursiveTree on the provided directory
	if err := recursiveTree(parentDir); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

