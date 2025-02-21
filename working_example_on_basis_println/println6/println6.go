package main

import "fmt"

func main() {
    // Create a variable called formatter, which is a string made of 4 formatting characters
    formatter := "%v %v %v %v"

    // Print formatter with 4 digits
    fmt.Printf(formatter+"\n", 1, 2, 3, 4)
    
    // Print formatter with 4 strings
    fmt.Printf(formatter+"\n", "one", "two", "three", "four")
    
    // Print formatter with 4 Booleans
    fmt.Printf(formatter+"\n", true, false, false, true)
    
    // Print formatter with itself, 4 times
    fmt.Printf(formatter+"\n", formatter, formatter, formatter, formatter)
    
    // Print formatter with 4 strings
    fmt.Printf(formatter+"\n", 
        "I had this thing.",
        "That you could type up right.",
        "But it didn't sing.",
        "So I said goodnight.",
    )
}

// 1 2 3 4
// one two three four
// true false false true
// %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v %v
// I had this thing. That you could type up right. But it didn't sing. So I said goodnight.
