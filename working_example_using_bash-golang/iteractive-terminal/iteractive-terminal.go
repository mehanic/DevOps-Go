package main

import (
    "fmt"
    "os"
)

// Function to check if the terminal is interactive
func isInteractive() bool {
    // Check if stdin, stdout, and stderr are connected to a terminal
    return isTerminal(os.Stdin.Fd()) && isTerminal(os.Stdout.Fd()) && isTerminal(os.Stderr.Fd())
}

// Helper function to check if a given file descriptor is a terminal
func isTerminal(fd uintptr) bool {
    fileInfo, err := os.Stdin.Stat()
    if err != nil {
        return false
    }
    return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

func main() {
    if isInteractive() {
        fmt.Println("Interactive shell")
        fmt.Printf("List of params: %v\n", os.Args)
    } else {
        fmt.Println("Non-interactive shell")
        fmt.Printf("List of params: %v\n", os.Args)
    }
}

