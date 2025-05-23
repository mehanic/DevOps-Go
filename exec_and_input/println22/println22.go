package main

import (
    "fmt"
)

// secretFormula calculates jelly beans, jars, and crates based on the starting point
func secretFormula(started int) (int, int, int) {
    jellyBeans := started * 500
    jars := jellyBeans / 1000
    crates := jars / 100
    return jellyBeans, jars, crates
}

func main() {
    fmt.Println("Let's practice everything.")
    fmt.Println("You'd need to know 'bout escapes with \\ that do \n newlines and \t tabs.")

    poem := `
    \tThe lovely world
    with logic so firmly planted
    cannot discern \n the needs of love
    nor comprehend passion from intuition
    and requires an explanation
    \n\t\twhere there is none.
    `

    fmt.Println("--------------")
    fmt.Println(poem)
    fmt.Println("--------------")

    five := 10 - 2 + 3 - 6
    fmt.Printf("This should be five: %d\n", five)

    startPoint := 10000
    beans, jars, crates := secretFormula(startPoint)

    fmt.Printf("With a starting point of: %d\n", startPoint)
    fmt.Printf("We'd have %d beans, %d jars, and %d crates.\n", beans, jars, crates)

    startPoint = startPoint / 10

    fmt.Println("We can also do that this way:")
    // Call secretFormula with the updated startPoint and print the results
    beans, jars, crates = secretFormula(startPoint)
    fmt.Printf("We'd have %d beans, %d jars, and %d crates.\n", beans, jars, crates)
}

