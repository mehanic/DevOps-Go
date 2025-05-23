package main

import "fmt"

// Declare a global variable
var globalVar int

// Function to print the value of the global variable
func myFunction() {
    fmt.Println(globalVar) // Print the value of the global variable
}

func main() {
    globalVar = 1 // Set the global variable

    // Call the function to print the global variable
    myFunction()
}

