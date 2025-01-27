package main

import (
	"fmt"
	"strings"
)

func main() {
	// The names of your three best friends
	names := [...]string{
		"Einstein",
		"Tesla",
		"Shepard",
	}

	// The distances to five different locations
	distances := [...]int{50, 40, 75, 30, 125}

	// A data buffer with five bytes of capacity
	data := [...]byte{'H', 'E', 'L', 'L', 'O'}

	// Currency exchange ratios only for a single currency
	ratios := [...]float64{3.14145}

	// Up/Down status of four different web servers
	alives := [...]bool{true, false, true, false}

	// A byte array that doesn't occupy memory space
	// Obviously, do not use ellipsis on this one
	var zero []byte

	// =========================================================================

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Print("\ndistances", separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndata", separator)
	for i := 0; i < len(data); i++ {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	fmt.Print("\nratios", separator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Print("\nalives", separator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}

	// =========================================================================

	// you know how this works :) don't be freaked out!
	fmt.Printf(`%sFOR RANGES%[1]s`, strings.Repeat("~", 30))

	fmt.Print("names", separator)
	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}

	fmt.Print("\ndistances", separator)
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	fmt.Print("\ndata", separator)
	for i, v := range data {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	fmt.Print("\nratios", separator)
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}

	fmt.Print("\nalives", separator)
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i, v := range zero {
		fmt.Printf("zero[%d]: %d\n", i, v)
	}
}

// names
// ====================
// names[0]: "Einstein"
// names[1]: "Tesla"
// names[2]: "Shepard"

// distances
// ====================
// distances[0]: 50
// distances[1]: 40
// distances[2]: 75
// distances[3]: 30
// distances[4]: 125

// data
// ====================
// data[0]: 72
// data[1]: 69
// data[2]: 76
// data[3]: 76
// data[4]: 79

// ratios
// ====================
// ratios[0]: 3.14

// alives
// ====================
// alives[0]: true
// alives[1]: false
// alives[2]: true
// alives[3]: false

// zero
// ====================
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~FOR RANGES~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~names
// ====================
// names[0]: "Einstein"
// names[1]: "Tesla"
// names[2]: "Shepard"

// distances
// ====================
// distances[0]: 50
// distances[1]: 40
// distances[2]: 75
// distances[3]: 30
// distances[4]: 125

// data
// ====================
// data[0]: 72
// data[1]: 69
// data[2]: 76
// data[3]: 76
// data[4]: 79

// ratios
// ====================
// ratios[0]: 3.14

// alives
// ====================
// alives[0]: true
// alives[1]: false
// alives[2]: true
// alives[3]: false

// zero
// ====================
