package main

import (
	"fmt"
	"math"
)

func main() {
	var adjacent, opposite float64

	// Prompt the user for the lengths
	fmt.Print("Enter the Adjacent length: ")
	fmt.Scanln(&adjacent) // Read adjacent length from input

	fmt.Print("Enter the Opposite length: ")
	fmt.Scanln(&opposite) // Read opposite length from input

	// Calculate the squares
	asquared := math.Pow(adjacent, 2) // a^2
	osquared := math.Pow(opposite, 2)  // o^2
	hsquared := asquared + osquared     // h^2 = a^2 + o^2

	// Calculate the hypotenuse
	hypotenuse := math.Sqrt(hsquared) // h = sqrt(h^2)

	// Print the result rounded to three decimal places
	fmt.Printf("The Hypotenuse is %.3f\n", hypotenuse)
}

