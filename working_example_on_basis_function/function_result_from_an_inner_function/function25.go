package main

import "fmt"

// myFunction takes an integer foo and returns the integer incremented by a result from an inner function
func myFunction(foo int) int {
    // Define an inner function myInnerFunction
    myInnerFunction := func() int {
        return 1
    }
    // Use the inner function and add its result to foo
    return foo + myInnerFunction()
}

func main() {
    // Call myFunction and print the result
    result := myFunction(1)
    fmt.Println(result) // Output will be 2
}

// 2
