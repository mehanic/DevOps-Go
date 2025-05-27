package main

import "fmt"

func lonelyInteger(nums []int) int {
    res := 0
    for _, num := range nums {
        res ^= num
    }
    return res
}

func main() {
    nums := []int{2, 3, 5, 4, 5, 3, 4}
    fmt.Println(lonelyInteger(nums)) // Output: 2
}
