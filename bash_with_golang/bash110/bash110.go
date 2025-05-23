package main

import "fmt"

func main() {
    // Print header
    fmt.Printf("%-5s %-10s %-4s\n", "no", "name", "mark")
    
    // Print formatted row
    fmt.Printf("%-5d %-10s %-4.2f\n", 1, "liyang", 10.00)
}

