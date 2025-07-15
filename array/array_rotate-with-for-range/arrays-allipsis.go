package main

import (
	"fmt"
	"strings"
)

func main() {
	separator := "\n" + strings.Repeat("=", 20) + "\n"
	fmt.Printf(`%sFOR RANGES%[1]s`, strings.Repeat("~", 30))

	// The names of your three best friends
	names := [...]string{
		"Einstein",
		"Tesla",
		"Shepard",
	}

	fmt.Print("names", separator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Print("names", separator)
	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}

	distances := [...]int{50, 40, 75, 30, 125}

	fmt.Print("\ndistances", separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndistances", separator)
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	// A data buffer with five bytes of capacity
	data := [...]byte{'H', 'E', 'L', 'L', 'O'}

	fmt.Print("\ndata", separator)
	for i := 0; i < len(data); i++ {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}
	fmt.Print("\ndata", separator)
	for i, v := range data {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, v)
	}
	fmt.Print("\ndata", separator)
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %c\n", i, data[i]) // Используем %c
	}

	fmt.Print("\ndata", separator)
	for i, v := range data {
		fmt.Printf("data[%d]: %c\n", i, v) // Используем %c
	}
	// Currency exchange ratios only for a single currency
	ratios := [...]float64{3.14145}

	fmt.Print("\nratios", separator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}
	fmt.Print("\nratios", separator)
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}

	// Up/Down status of four different web servers
	alives := [...]bool{true, false, true, false}

	fmt.Print("\nalives", separator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}
	fmt.Print("\nalives", separator)
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}

	// A byte array that doesn't occupy memory space
	// Obviously, do not use ellipsis on this one
	var zero []byte

	fmt.Print("\nzero", separator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}
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
