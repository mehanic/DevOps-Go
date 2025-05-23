package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    dir := "/tmp"

    // Read the directory
    files, err := os.ReadDir(dir)
    if err != nil {
        fmt.Printf("Error reading directory %s: %v\n", dir, err)
        return
    }

    // Check each file in the directory
    for _, file := range files {
        isExecutable(filepath.Join(dir, file.Name()))
    }
}

// Function to check if a file is executable
func isExecutable(file string) {
    info, err := os.Stat(file)
    if err != nil {
        fmt.Printf("Error stating file %s: %v\n", file, err)
        return
    }

    // Check if the file is executable
    if info.Mode().Perm()&0111 != 0 { // Check if executable bit is set
        fmt.Printf("%s is executable.\n", file)
    } else {
        fmt.Printf("%s is not executable.\n", file)
    }
}

