package main

import (
    "fmt"
    "os"
)

func main() {
    var n uint
    fmt.Print("Input a number: ")
    if _, err := fmt.Scan(&n); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
    
    var sum uint
    for i := uint(1); i < n; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += i
        }
    }
    fmt.Println("Sum of 3|5 multiples below:", sum)
}
