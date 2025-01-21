package main

import (
    "fmt"
)

func main() {
    for each := 1; each <= 10; each++ {
        if each == 7 {
            continue
        }
        fmt.Println("this is the line inside of your if condition after continue keyword")
        fmt.Println(each)
    }
}

// this is the line inside of your if condition after continue keyword
// 1
// this is the line inside of your if condition after continue keyword
// 2
// this is the line inside of your if condition after continue keyword
// 3
// this is the line inside of your if condition after continue keyword
// 4
// this is the line inside of your if condition after continue keyword
// 5
// this is the line inside of your if condition after continue keyword
// 6
// this is the line inside of your if condition after continue keyword
// 8
// this is the line inside of your if condition after continue keyword
// 9
// this is the line inside of your if condition after continue keyword
// 10
