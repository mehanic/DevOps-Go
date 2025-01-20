package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	startTime := time.Now()

	// Calculate the product of the first 100,000 numbers
	product := new(big.Int).SetInt64(1)
	for i := int64(1); i < 100000; i++ {
		product.Mul(product, big.NewInt(i))
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	// Convert product to string and get its length
	productStr := product.String()
	numDigits := len(productStr)

	fmt.Printf("The result is %d digits long.\n", numDigits)
	fmt.Printf("Took %s seconds to calculate.\n", duration)
}

// The result is 456569 digits long.
// Took 838.606462ms seconds to calculate.
