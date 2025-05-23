package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func printAll(file *os.File) {
    // Create a scanner to read the file
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}

func rewind(file *os.File) {
    // Move the file pointer back to the beginning
    file.Seek(0, 0)
}

func printALine(lineCount int, file *os.File) {
    // Create a scanner to read the file
    scanner := bufio.NewScanner(file)
    for i := 1; i <= lineCount; i++ {
        if scanner.Scan() {
            if i == lineCount {
                fmt.Println(lineCount, scanner.Text())
                return
            }
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run script.go <filename>")
        return
    }

    script := os.Args[0]
    inputFile := os.Args[1]

    fmt.Printf("Opening %s...\n", inputFile)
    file, err := os.Open(inputFile)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    fmt.Println("First let's print the whole file:\n")
    printAll(file)

    fmt.Println("Now let's rewind, kind of like a tape.")
    rewind(file)

    fmt.Println("Let's print three lines:")

    currentLine := 1
    printALine(currentLine, file)

    currentLine++
    printALine(currentLine, file)

    currentLine++
    printALine(currentLine, file)
}

