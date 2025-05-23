package main

import "fmt"

// addNumbers adds numbers from 0 to max (exclusive) with a given step and returns a slice of those numbers
func addNumbers(max, step int) []int {
    i := 0
    var numbers []int

    for i < max {
        fmt.Printf("At the top i is %d\n", i)
        numbers = append(numbers, i)

        i = i + step
        fmt.Printf("Numbers now: %v\n", numbers)
        fmt.Printf("At the bottom i is %d\n", i)
    }

    return numbers
}

func main() {
    numbers := addNumbers(10, 5)

    fmt.Println("The numbers: ")
    for _, num := range numbers {
        fmt.Println(num)
    }
}

// At the top i is 0
// Numbers now: [0]
// At the bottom i is 5
// At the top i is 5
// Numbers now: [0 5]
// At the bottom i is 10
// The numbers: 
// 0
// 5
