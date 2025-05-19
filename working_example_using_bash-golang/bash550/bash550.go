package main

import (
    "fmt"
    "os"
)

func main() {
    // Check if at least 3 arguments are provided
    if len(os.Args) < 4 {
        fmt.Println("Usage: go run main.go <arg1> <arg2> <arg3>")
        return
    }

    // Print first three command line arguments
    fmt.Println(os.Args[1], os.Args[2], os.Args[3], "-> echo $1 $2 $3")

    // Store all arguments in a slice
    args := os.Args[1:] // Store from index 1 to exclude the program name
    fmt.Println(args[0], args[1], args[2], "-> args=($@); echo ${args[0]} ${args[1]} ${args[2]}")

    // Print all arguments at once
    fmt.Println(os.Args[1:], "-> $@")

    // Print the number of arguments passed to the program
    fmt.Println("Number of arguments passed:", len(os.Args)-1, "-> echo Number of arguments passed: $#")
}

