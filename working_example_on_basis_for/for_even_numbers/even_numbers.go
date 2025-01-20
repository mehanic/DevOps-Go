package main

import (
	"fmt"
)

func main() {
	// Initialize a slice to hold the even numbers
	var evenNumbers []int
	
	// Populate the slice with even numbers from 2 to 10
	for i := 2; i < 11; i += 2 {
		evenNumbers = append(evenNumbers, i)
	}

	// Print the slice
	fmt.Println(evenNumbers)
}


//[2 4 6 8 10]
