package main

import (
    "fmt"
)

func hammingWeight(n int) int {
    res := 0
    for i := 0; i < 32; i++ {
        if (1<<i)&n != 0 {
            res++
        }
    }
    return res
}

func main() {
    fmt.Println(hammingWeight(11))  // 11 в двоичном: 1011, количество единиц = 3
    fmt.Println(hammingWeight(128)) // 128 в двоичном: 10000000, количество единиц = 1
    fmt.Println(hammingWeight(0))   // 0 в двоичном: 0, количество единиц = 0
}
