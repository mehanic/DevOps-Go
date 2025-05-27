package main

import "fmt"

func triangleNumbers(n int) int {
    if n%2 != 0 {
        return 2
    } else if n%4 == 0 {
        return 3
    }
    return 4
}

func main() {
    fmt.Println(triangleNumbers(1))  // 2
    fmt.Println(triangleNumbers(4))  // 3
    fmt.Println(triangleNumbers(6))  // 4
}
