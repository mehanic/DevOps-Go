package main

import (
    "fmt"
    "math/big"
    "time"
)

func main() {
    // Record the start time
    startTime := time.Now()

    // Calculate the product of the first 100,000 numbers
    product := big.NewInt(1)
    for i := 1; i < 100000; i++ {
        product.Mul(product, big.NewInt(int64(i)))
    }

    // Record the end time
    endTime := time.Now()

    // Calculate the duration
    duration := endTime.Sub(startTime)

    // Get the length of the result in digits
    numDigits := len(product.String())

    // Print the results
    fmt.Printf("The result is %d digits long.\n", numDigits)
    fmt.Printf("Took %s seconds to calculate.\n", duration)
}

