package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
//	"strings"
)

func main() {
	// Step 1: Create a file named log.1
	err := ioutil.WriteFile("log.1", []byte("This is a log file."), 0644)
	if err != nil {
		fmt.Println("Error creating log.1:", err)
		return
	}

	// Step 2: Create a symbolic link named log pointing to log.1
	err = os.Symlink("log.1", "log")
	if err != nil {
		fmt.Println("Error creating symbolic link log:", err)
		return
	}

	// Step 3: List symbolic links in the current directory
	fmt.Println("Symbolic links in the current directory:")

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		// Check if the file is a symbolic link
		if file.Mode()&os.ModeSymlink != 0 {
			fmt.Println(file.Name())
		}
	}

	// Step 4: Find all symbolic links in the current directory and its subdirectories
	fmt.Println("\nFinding all symbolic links recursively:")
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the file is a symbolic link
		if info.Mode()&os.ModeSymlink != 0 {
			fmt.Println(path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through the directory:", err)
	}
}

