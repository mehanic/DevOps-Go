package main

import "fmt"

func main() {
    numbers := []int{3, 4, 56, 7, 8}
    for _, each := range numbers {
        fmt.Println(each)
        if each == 56 {
            break
        }
    }

    fmt.Println("after loop")
}

// 3
// 4
// 56
// after loop
