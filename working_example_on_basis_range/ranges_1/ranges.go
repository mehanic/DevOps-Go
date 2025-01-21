package main

import "fmt"

// Rangess prints the names of cities using two different methods.
func Rangess() {
    sehirler := []string{"Ankara", "Izmir", "Istanbul"}

    // First method: using a traditional for loop with an index
    fmt.Println("Using traditional for loop:")
    for i := 0; i < len(sehirler); i++ {
        fmt.Println(sehirler[i])
    }

    // Second method: using a range-based for loop
    fmt.Println("\nUsing range-based for loop:")
    for _, sehir := range sehirler {
        fmt.Println(sehir)
    }
}

func main() {
    Rangess()
}


// Using traditional for loop:
// Ankara
// Izmir
// Istanbul

// Using range-based for loop:
// Ankara
// Izmir
// Istanbul