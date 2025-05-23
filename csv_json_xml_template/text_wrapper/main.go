package main

import (
	"fmt"
	"unicode"
)

func main() {
	const text = `At one end of the Western Spiral Arm of the galaxy, in a remote and uncharted corner, there is a small, yellow sun, hidden from sight.

Orbiting this sun, roughly one hundred and forty-eight million kilometers away, is a completely insignificant, blue-green little planet.

The planet's inhabitants, who are descended from apes, are so primitive that they still consider digital wristwatches to be a remarkably impressive invention.`

	const maxWidth = 40

	var lw int // line width

	for _, r := range text {
		fmt.Printf("%c", r)

		switch lw++; {
		case lw > maxWidth && r != '\n' && unicode.IsSpace(r):
			fmt.Println()
			fallthrough
		case r == '\n':
			lw = 0
		}
	}
	fmt.Println()
}



// At one end of the Western Spiral Arm of the 
// galaxy, in a remote and uncharted corner, 
// there is a small, yellow sun, hidden from 
// sight.

// Orbiting this sun, roughly one hundred and 
// forty-eight million kilometers away, is a 
// completely insignificant, blue-green little 
// planet.

// The planet's inhabitants, who are descended 
// from apes, are so primitive that they still 
// consider digital wristwatches to be a remarkably 
// impressive invention.
