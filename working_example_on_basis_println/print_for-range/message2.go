package main

import (
	"errors"
	"fmt"
	"strings"
)

func boxPrint(symbol string, width, height int) error {
	if len(symbol) != 1 {
		return errors.New("symbol must be a single character string")
	}
	if width <= 2 {
		return errors.New("width must be greater than 2")
	}
	if height <= 2 {
		return errors.New("height must be greater than 2")
	}

	fmt.Println(strings.Repeat(symbol, width))
	for i := 0; i < height-2; i++ {
		fmt.Println(symbol + strings.Repeat(" ", width-2) + symbol)
	}
	fmt.Println(strings.Repeat(symbol, width))
	return nil
}

func main() {
	testCases := [][3]interface{}{
		{'*', 4, 4},  // rune for '*'
		{'O', 20, 5}, // rune for 'O'
		{'x', 1, 3},  // rune for 'x'
		{'Z', 3, 3},  // rune for 'Z'
	}

	for _, tc := range testCases {
		// Convert rune to string
		sym := string(tc[0].(rune)) // Convert rune to string
		w := tc[1].(int)
		h := tc[2].(int)
		err := boxPrint(sym, w, h)
		if err != nil {
			fmt.Println("An exception happened:", err)
		}
	}
}


// ****
// *  *
// *  *
// ****
// OOOOOOOOOOOOOOOOOOOO
// O                  O
// O                  O
// O                  O
// OOOOOOOOOOOOOOOOOOOO
// An exception happened: width must be greater than 2
// ZZZ
// Z Z
// ZZZ
