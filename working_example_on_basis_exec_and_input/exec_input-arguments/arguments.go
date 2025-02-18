package main

import (
    "fmt"
)

// printTwo takes two arguments and prints them.
func printTwo(arg1, arg2 string) {
    fmt.Printf("arg1: %s, arg2: %s\n", arg1, arg2)
}

// printTwoAgain takes two arguments and prints them.
func printTwoAgain(arg1, arg2 string) {
    fmt.Printf("arg1: %s, arg2: %s\n", arg1, arg2)
}

// printOne takes one argument and prints it.
func printOne(arg1 string) {
    fmt.Printf("arg1: %s\n", arg1)
}

// printNone prints a default message with no arguments.
func printNone() {
    fmt.Println("I got nothin'.")
}

func main() {
    printTwo("Zed", "Shaw")
    printTwoAgain("Zed", "Shaw")
    printOne("First!")
    printNone()
}

// arg1: Zed, arg2: Shaw
// arg1: Zed, arg2: Shaw
// arg1: First!
// I got nothin'.
