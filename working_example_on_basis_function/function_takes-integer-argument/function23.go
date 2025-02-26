package main

import "fmt"

// myFunction takes an integer argument and returns the integer incremented by 1
func myFunction(myArg int) int {
    return myArg + 1
}

func main() {
    result := myFunction(5)
    fmt.Println(result) // Output will be 6
}

// 6

