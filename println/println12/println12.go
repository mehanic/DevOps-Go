package main

import (
    "fmt"
    "os"
)

func main() {
    // Check if enough command-line arguments are provided
    if len(os.Args) != 4 {
        fmt.Println("Usage: go run script.go <first> <second> <third>")
        os.Exit(1)
    }

    // Assign command-line arguments to variables
    script := os.Args[0]
    first := os.Args[1]
    second := os.Args[2]
    third := os.Args[3]

    // Print the variables
    fmt.Println("The script is called:", script)
    fmt.Println("The first variable is:", first)
    fmt.Println("The second variable is:", second)
    fmt.Println("Your third variable is:", third)
}

// The script is called: /tmp/go-build3592191095/b001/exe/println12
// The first variable is: 1
// The second variable is: 2
// Your third variable is: 3
