package main

import "fmt"

func main() {
	const word = "console"

	for _, w := range word {
		fmt.Printf("%c\n", w)
		fmt.Printf("\tdecimal: %[1]d\n", w)
		fmt.Printf("\thex    : 0x%[1]x\n", w)
		fmt.Printf("\tbinary : 0b%08[1]b\n", w)
	}
	fmt.Println("---------------")
	// print the word manually using runes
	fmt.Printf("with runes       : %s\n",
		string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))

	// print the word manually using decimals
	fmt.Printf("with decimals    : %s\n",
		string([]byte{99, 111, 110, 115, 111, 108, 101}))

	// print the word manually using hexadecimals
	fmt.Printf("with hexadecimals: %s\n",
		string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))
}

// c
// 	decimal: 99
// 	hex    : 0x63
// 	binary : 0b01100011
// o
// 	decimal: 111
// 	hex    : 0x6f
// 	binary : 0b01101111
// n
// 	decimal: 110
// 	hex    : 0x6e
// 	binary : 0b01101110
// s
// 	decimal: 115
// 	hex    : 0x73
// 	binary : 0b01110011
// o
// 	decimal: 111
// 	hex    : 0x6f
// 	binary : 0b01101111
// l
// 	decimal: 108
// 	hex    : 0x6c
// 	binary : 0b01101100
// e
// 	decimal: 101
// 	hex    : 0x65
// 	binary : 0b01100101
// ---------------
// with runes       : console
// with decimals    : console
// with hexadecimals: console
