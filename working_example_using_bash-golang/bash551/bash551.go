package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("[*] Usage: go run main.go domains.txt")
        return
    }

    domainFile := os.Args[1]
    
    // Check if the file exists
    if _, err := os.Stat(domainFile); os.IsNotExist(err) {
        fmt.Println("[*] Usage: go run main.go domains.txt")
        return
    }

    // Open the file
    file, err := os.Open(domainFile)
    if err != nil {
        fmt.Printf("Error opening file: %s\n", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        domain := scanner.Text()
        // Resolve the domain using the host command
        output, err := exec.Command("host", domain).Output()
        if err != nil {
            fmt.Printf("Error resolving domain %s: %s\n", domain, err)
            continue
        }

        // Process the output
        lines := strings.Split(string(output), "\n")
        for _, line := range lines {
            if strings.Contains(line, "has address") {
                parts := strings.Split(line, " has address ")
                if len(parts) == 2 {
                    fmt.Printf("%s : %s\n", domain, parts[1])
                }
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %s\n", err)
    }
}

