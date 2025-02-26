package main

import "fmt"

// myFunction takes an integer and returns the integer incremented by 1
func myFunction(foo int) int {
    return foo + 1
}

func main() {
    // Assign the function to a variable
    var myVar = myFunction

    // Call the function through the variable and print the result
    result := myVar(1)
    fmt.Println(result) // Output will be 2
}

// 2
