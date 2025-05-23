package main

import (
	"fmt"
	"math"
)

func main() {
	// Get floating point number
	floatingPointNumber := 3.3446
	fmt.Println(floatingPointNumber)

	// Round floating point number
	bashRoundedNumber := math.Round(floatingPointNumber)
	fmt.Println("Rounded number with Go:", bashRoundedNumber)
}

