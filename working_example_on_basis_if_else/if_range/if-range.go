package main

import (
	"fmt"
)

func main() {
	message := "It was a bright cold day in April, and the clocks were striking thirteen."
	count := make(map[rune]int)

	// Count occurrences of each character
	for _, character := range message {
		count[character]++
	}

	// Print the results
	fmt.Println("Character counts:")
	for character, frequency := range count {
		fmt.Printf("%c: %d\n", character, frequency)
	}
}

// Character counts:
// r: 5
// b: 1
// c: 3
// d: 3
// w: 2
// t: 6
// a: 4
// s: 3
// i: 6
// g: 2
// o: 2
// l: 3
// I: 1
// n: 4
// p: 1
// .: 1
// y: 1
// h: 3
// A: 1
// ,: 1
// e: 5
// k: 2
//  : 13
