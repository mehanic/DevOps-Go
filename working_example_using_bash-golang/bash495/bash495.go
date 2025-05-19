package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Since we're dealing with paths, set current working directory.
	// Not strictly necessary in Go as the working directory is already set when executing.
	// If you need to change it, you can use os.Chdir.

	// Input validation.
	if len(os.Args) != 2 {
		fmt.Println("Incorrect usage!")
		fmt.Println("Usage: go run print-or-list.go <file or directory path>")
		os.Exit(1)
	}

	inputPath := os.Args[1]

	// Check if the path is a file
	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		log.Fatal("Cannot access path:", err)
	}

	if fileInfo.Mode().IsRegular() {
		// It's a file, print the content
		fmt.Println("File found, showing content:")
		content, err := ioutil.ReadFile(inputPath)
		if err != nil {
			fmt.Println("Cannot print file, exiting program!")
			os.Exit(1)
		}
		fmt.Println(string(content))
	} else if fileInfo.IsDir() {
		// It's a directory, list its contents
		fmt.Println("Directory found, listing:")
		files, err := ioutil.ReadDir(inputPath)
		if err != nil {
			fmt.Println("Cannot list directory, exiting program!")
			os.Exit(1)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	} else {
		// It's neither a file nor a directory
		fmt.Println("Path is neither a file nor a directory, exiting program.")
		os.Exit(1)
	}
}

