package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
//	"path/filepath"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 2 {
		fmt.Println("Incorrect usage!")
		fmt.Println("Usage: go run nested-print-or-list.go <file or directory path>")
		os.Exit(1)
	}

	inputPath := os.Args[1]

	// Check if the path is readable
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fmt.Println("Cannot read the file/directory, exiting program.")
		os.Exit(1)
	}

	// Determine if the path is a file or directory
	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	if fileInfo.Mode().IsRegular() {
		// It's a file, print the content
		fmt.Println("File found, showing content:")
		content, err := ioutil.ReadFile(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(content))
	} else if fileInfo.IsDir() {
		// It's a directory, list its contents
		fmt.Println("Directory found, listing:")
		files, err := ioutil.ReadDir(inputPath)
		if err != nil {
			log.Fatal(err)
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

