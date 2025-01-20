package main

import (
	"fmt"
)

func main() {
	// Declare and initialize the slice
	myPets := []string{"Zophie", "Pooka", "Fat-tail"}
	fmt.Println(myPets)

	// Assign individual elements of the slice to variables
	a, b, c := myPets[0], myPets[1], myPets[2]

	// Iterate over the individual variables and print them
	for _, pet := range []string{a, b, c} {
		fmt.Println(pet)
	}
}

// [Zophie Pooka Fat-tail]
// Zophie
// Pooka
// Fat-tail
