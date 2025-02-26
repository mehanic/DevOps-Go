package main

import "fmt"

// Define a function type that takes an int and returns an int
type IntFunc func(int) int

// myFunction returns a function of type IntFunc
func myFunction(constant int) IntFunc {
    // Define and return the inner function
    return func(foo int) int {
        return foo + constant
    }
}

func main() {
    // Get a function that adds 1
    plusOne := myFunction(1)
    // Call the returned function with the argument 1
    result := plusOne(1)
    fmt.Println(result) // Output will be 2
}

// 2
