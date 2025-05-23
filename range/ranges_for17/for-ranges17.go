package main

import "fmt"

func main() {
    // Define the name variable
    name := "Student"

    // Print the name
    fmt.Printf("My name is %s\n", name)

    // Use a loop to print the name multiple times
    i := 0
    for i < 5 {
        fmt.Printf("%s five times (%d)\n", name, i)
        i++
    }
}

// My name is Student
// Student five times (0)
// Student five times (1)
// Student five times (2)
// Student five times (3)
// Student five times (4)
