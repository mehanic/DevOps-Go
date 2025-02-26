package main

import "fmt"

// Define the type for a function that returns an int
type IntFunc func() int

// myFunction takes an integer foo and a function as a parameter
func myFunction(foo int, myParameterFunction IntFunc) int {
    return foo + myParameterFunction()
}

// parameterFunction returns 1
func parameterFunction() int {
    return 1
}

func main() {
    // Call myFunction with 1 and parameterFunction as arguments
    result := myFunction(1, parameterFunction)
    fmt.Println(result) // Output will be 2
}

// 2
