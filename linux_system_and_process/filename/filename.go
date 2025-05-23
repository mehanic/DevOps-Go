package main

import (
    "fmt"
    "os"
)

func main() {
    // Ensure that exactly 2 command-line arguments are provided (including the script name)
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run script.go <filename>")
        os.Exit(1)
    }

    filename := os.Args[1]

    fmt.Printf("Opening %s...\n", filename)

    // Read the contents of the file
    content, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    fmt.Printf("Printing the contents of %s...\n", filename)
    fmt.Println(string(content))
}

