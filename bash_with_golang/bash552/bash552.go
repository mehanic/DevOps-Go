package main

import (
    "fmt"
    "os"
//    "os/user"
)

func main() {
    // Get the current script's filename
    file := os.Args[0]

    // Check if the file exists
    if _, err := os.Stat(file); err == nil {
        fmt.Printf("%s exists\n", file)
    } else {
        fmt.Printf("%s does not exist\n", file)
        return
    }

    // Check if the file is readable
    if isReadable(file) {
        fmt.Printf("%s can be read\n", file)
    } else {
        fmt.Printf("%s cannot be read\n", file)
    }

    // Check if the file is writable
    if isWritable(file) {
        fmt.Printf("%s can be written to\n", file)
    } else {
        fmt.Printf("%s cannot be written to\n", file)
    }

    // Check if the file is executable
    if isExecutable(file) {
        fmt.Printf("%s can be executed\n", file)
    } else {
        fmt.Printf("%s cannot be executed\n", file)
    }
}

// Function to check if the file is readable
func isReadable(file string) bool {
    f, err := os.Open(file)
    if err != nil {
        return false
    }
    defer f.Close()
    return true
}

// Function to check if the file is writable
func isWritable(file string) bool {
    f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return false
    }
    defer f.Close()
    return true
}

// Function to check if the file is executable
func isExecutable(file string) bool {
    info, err := os.Stat(file)
    if err != nil {
        return false
    }

    return info.Mode().Perm()&0111 != 0 // Check if executable bit is set
}

