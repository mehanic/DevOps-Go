package main

import (
    "fmt"
    "os"
)

func main() {
    // Check if exactly 3 command-line arguments are provided (including the script name)
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run script.go <first_name> <last_name>")
        os.Exit(1)
    }

    // Assign command-line arguments to variables
    firstName := os.Args[1]
    lastName := os.Args[2]

    // Print the full name
    fmt.Printf("Your full name is %s %s\n", firstName, lastName)
}

