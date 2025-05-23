package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
)

func main() {
    // Ensure that exactly 2 command-line arguments are provided (including the script name)
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run script.go <filename>")
        os.Exit(1)
    }

    // Assign command-line arguments to variables
    filename := os.Args[1]

    // Open the file
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Print the filename
    fmt.Printf("Here's your file %s:\n", filename)

    // Read and print the file contents
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // Print a string
    fmt.Println("I'll also ask you to type it again:")

    // Prompt the user for a file name
    fmt.Print("> ")
    var fileAgain string
    fmt.Scanln(&fileAgain)

    // Open the file that the user entered
    fileAgainHandle, err := os.Open(fileAgain)
    if err != nil {
        log.Fatal(err)
    }
    defer fileAgainHandle.Close()

    // Read and print the file contents
    scanner = bufio.NewScanner(fileAgainHandle)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

