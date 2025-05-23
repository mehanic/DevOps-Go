package main

import "fmt"

// cheeseAndCrackers prints out information about cheese and crackers.
func cheeseAndCrackers(cheeseCount, boxesOfCrackers int) {
    fmt.Printf("You have %d cheeses!\n", cheeseCount)
    fmt.Printf("You have %d boxes of crackers!\n", boxesOfCrackers)
    fmt.Println("Man that's enough for a party!")
    fmt.Println("Get a blanket.\n")
}

func main() {
    // Print a string
    fmt.Println("We can just give the function numbers directly:")
    // Pass 20, 30 to the cheeseAndCrackers function
    cheeseAndCrackers(20, 30)

    // Print a string
    fmt.Println("OR, we can use variables from our script:")
    // Set a variable to 10
    amountOfCheese := 10
    // Set a variable to 50
    amountOfCrackers := 50

    // Pass amountOfCheese and amountOfCrackers variables to cheeseAndCrackers function
    cheeseAndCrackers(amountOfCheese, amountOfCrackers)

    // Print a string
    fmt.Println("We can even do math inside too:")
    // Pass 10 + 20 and 5 + 6 to the cheeseAndCrackers function
    cheeseAndCrackers(10 + 20, 5 + 6)

    // Print a string
    fmt.Println("And we can combine the two, variables and math:")
    // Pass the amountOfCheese variable + 100 and the amountOfCrackers variable + 1000 to cheeseAndCrackers function
    cheeseAndCrackers(amountOfCheese + 100, amountOfCrackers + 1000)
}

