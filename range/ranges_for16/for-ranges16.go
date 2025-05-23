package main

import "fmt"

func main() {
    // Define the name variable
    name := "Student"

    // Print the name
    fmt.Printf("My name is %s\n", name)

    // Use a loop to print the name multiple times
    for i := 0; i < 5; i++ {
        fmt.Printf("%s Five Times (%d)\n", name, i)
    }
}

// My name is Student
// Student Five Times (0)
// Student Five Times (1)
// Student Five Times (2)
// Student Five Times (3)
// Student Five Times (4)
