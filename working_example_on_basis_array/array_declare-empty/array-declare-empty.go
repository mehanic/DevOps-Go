package main

import (
	"fmt"
)

// declare empty
func main() {
	var (
		names     [3]string  // The names of your three best friends
		distances [5]int     // The distances to five different locations
		data      [5]byte    // A data buffer with five bytes of capacity
		ratios    [1]float64 // Currency exchange ratios only for a single currency
		alives    [4]bool    // Up/Down status of four different web servers
		zero      [0]byte    // A byte array that doesn't occupy memory space
	)

	// 1. Declare and print the arrays with their types.
	fmt.Printf("names    : %#v\n", names)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data     : %#v\n", data)
	fmt.Printf("ratios   : %#v\n", ratios)
	fmt.Printf("alives   : %#v\n", alives)
	fmt.Printf("zero     : %#v\n", zero)

	//  2. Print only the types of the same arrays.
	fmt.Println()
	fmt.Printf("names    : %T\n", names)
	fmt.Printf("distances: %T\n", distances)
	fmt.Printf("data     : %T\n", data)
	fmt.Printf("ratios   : %T\n", ratios)
	fmt.Printf("alives   : %T\n", alives)
	fmt.Printf("zero     : %T\n", zero)

	// 3. Print only the elements of the same arrays.
	fmt.Println()
	fmt.Printf("names    : %q\n", names)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("data     : %d\n", data)
	fmt.Printf("ratios   : %.2f\n", ratios)
	fmt.Printf("alives   : %t\n", alives)
	fmt.Printf("zero     : %d\n", zero)
}

// names    : [3]string{"", "", ""}
// distances: [5]int{0, 0, 0, 0, 0}
// data     : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
// ratios   : [1]float64{0}
// alives   : [4]bool{false, false, false, false}
// zero     : [0]uint8{}

// names    : [3]string
// distances: [5]int
// data     : [5]uint8
// ratios   : [1]float64
// alives   : [4]bool
// zero     : [0]uint8

// names    : ["" "" ""]
// distances: [0 0 0 0 0]
// data     : [0 0 0 0 0]
// ratios   : [0.00]
// alives   : [false false false false]
// zero     : []
