package main

import "fmt"

// MakePizza function takes a size and a variable number of toppings
func MakePizza(size int, toppings ...string) {
    fmt.Printf("Making a %d-inch pizza with %v\n", size, toppings)
}

func main() {
    MakePizza(16, "pepperoni")
    MakePizza(12, "mushrooms", "green peppers", "extra cheese")
}

// Making a 16-inch pizza with [pepperoni]
// Making a 12-inch pizza with [mushrooms green peppers extra cheese]
