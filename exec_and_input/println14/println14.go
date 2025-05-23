package main

import (
    "fmt"
    "os"
)

func main() {
    // Check if exactly 5 command-line arguments are provided (including the script name)
    if len(os.Args) != 5 {
        fmt.Println("Usage: go run script.go <num1> <num2> <num3> <num4>")
        os.Exit(1)
    }

    // Assign command-line arguments to variables
    num1 := os.Args[1]
    num2 := os.Args[2]
    num3 := os.Args[3]
    num4 := os.Args[4]

    // Print the numbers
    fmt.Printf("The four numbers you entered are: %s, %s, %s and %s\n", num1, num2, num3, num4)
}

//The four numbers you entered are: 2, 3, 4 and 5