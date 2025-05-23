package main

import "fmt"

// add function takes two integers, prints the operation, and returns their sum
func add(a, b int) int {
    fmt.Printf("ADDING %d + %d\n", a, b)
    return a + b
}

// subtract function takes two integers, prints the operation, and returns their difference
func subtract(a, b int) int {
    fmt.Printf("SUBTRACTING %d - %d\n", a, b)
    return a - b
}

// multiply function takes two integers, prints the operation, and returns their product
func multiply(a, b int) int {
    fmt.Printf("MULTIPLYING %d * %d\n", a, b)
    return a * b
}

// divide function takes two integers, prints the operation, and returns their quotient
func divide(a, b int) int {
    fmt.Printf("DIVIDING %d / %d\n", a, b)
    return a / b
}

func main() {
    fmt.Println("Let's do some math with just functions!")

    age := add(30, 5)
    height := subtract(78, 4)
    weight := multiply(90, 2)
    iq := divide(100, 2)

    fmt.Printf("Age: %d, Height: %d, Weight: %d, IQ: %d\n", age, height, weight, iq)

    fmt.Println("Here is a puzzle.")

    // Calculate age + height - weight * (iq / 2)
    // 35  + 74     - 180    * (50 / 2)
    what := add(age, subtract(height, multiply(weight, divide(iq, 4))))

    fmt.Printf("That becomes: %d Can you do it by hand?\n", what)

    // Calculate 25 + (75 + (125 * 3))
    what2 := add(25, add(75, multiply(125, 3)))
    fmt.Printf("That becomes: %d\n", what2)
}

