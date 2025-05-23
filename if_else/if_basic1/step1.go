package main

import "fmt"

func main() {

    a := 5
    d := 8

    if a == d {
        fmt.Println("Both are equal")
    } else if a > d {
        fmt.Println("a is greater than d")
    } else if d > a {
        fmt.Println("a is less than d")
    } else {
        fmt.Println("You didn’t use your head")
    }
}
