package main

import "fmt"

func main() {
    // Lists (slices in Go)
    theCount := []int{1, 2, 3, 4, 5}
    fruits := []string{"apples", "oranges", "pears", "apricots"}
    change := []interface{}{1, "pennies", 2, "dimes", 3, "quarters"}

    // Iterating through a slice
    for _, number := range theCount {
        fmt.Printf("This is count %d\n", number)
    }

    // Iterating through another slice
    for _, fruit := range fruits {
        fmt.Printf("A fruit of type: %s\n", fruit)
    }

    // Iterating through a mixed slice
    for _, i := range change {
        fmt.Printf("I got %v\n", i) // Using %v for default formatting
    }

    // Building a new slice
    elements := make([]int, 0)
    for i := 0; i <= 5; i++ {
        elements = append(elements, i)
    }

    // Printing elements
    for _, i := range elements {
        fmt.Printf("Element was: %d\n", i)
    }
}

// This is count 1
// This is count 2
// This is count 3
// This is count 4
// This is count 5
// A fruit of type: apples
// A fruit of type: oranges
// A fruit of type: pears
// A fruit of type: apricots
// I got 1
// I got pennies
// I got 2
// I got dimes
// I got 3
// I got quarters
// Element was: 0
// Element was: 1
// Element was: 2
// Element was: 3
// Element was: 4
// Element was: 5
