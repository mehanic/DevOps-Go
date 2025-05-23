package main

import (
	"fmt"
	"os"
)



func printDirectoryContents(directory string) {
    dirEntries, err := os.ReadDir(directory)
    if err != nil {
        fmt.Println("Error reading directory:", err)
        return
    }

    for _, entry := range dirEntries {
        fmt.Printf("Name: %s\n", entry.Name())
        fmt.Printf("Is Directory: %t\n", entry.IsDir())
        info, err := entry.Info()
        if err == nil {
            fmt.Printf("Size: %d bytes\n", info.Size())
            fmt.Printf("Mode: %s\n", info.Mode())
            fmt.Printf("ModTime: %s\n", info.ModTime())
        } else {
            fmt.Println("Error getting file info:", err)
        }
        fmt.Println()
    }
}

func main() {
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
    //directory := "images/" // You can change this to any directory path you want to print
	// directory := os.ReadDir("images/")
	// directory := os.UserHomeDir()

    fmt.Printf("Contents of directory: %s\n\n", directory)
    printDirectoryContents(directory)
}

// Name: directory2.go
// Is Directory: false
// Size: 1098 bytes
// Mode: -rw-rw-r--
// ModTime: 2025-01-17 22:01:52.145885402 +0100 CET

