package main

import (
    "fmt"
)

func main() {
    count := 0
    Test := "Test"

    // Simulate: test $Test = 'Test' && echo True
    if Test == "Test" {
        fmt.Println("True")
    }

    // Simulate: [ $count -ne 0 ] || echo True
    if count != 0 {
        // Do nothing if count is not zero
    } else {
        fmt.Println("True")
    }

    // if-else equivalent for the Bash block
    if count != 0 {
        fmt.Println("false")
    } else if Test == "Test" && count == 0 {
        fmt.Println("latter is True")
    } else {
        fmt.Println("Both are false")
    }
}

