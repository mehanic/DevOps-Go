package main

import "fmt"

// Declare a global variable
var globalVar int

// Function to modify the global variable
func myFunction() {
    globalVar = 1 // Set the global variable
}

func main() {
    // Output the value of the global variable before calling the function
    fmt.Printf("GLOBAL_VAR value BEFORE myFunction called: %d\n", globalVar)

    // Call the function to set the global variable
    myFunction()

    // Output the value of the global variable after calling the function
    fmt.Printf("GLOBAL_VAR value AFTER myFunction called: %d\n", globalVar)
}

