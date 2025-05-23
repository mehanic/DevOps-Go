package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%d %b %x %X \n", 44, 28, 78996633, 44886347899) // % x for base 16 %X is hexdecimal with upper case
	fmt.Printf("%x \n", 29939)
	fmt.Printf("%#x \n", 29939) // # will add a leading 0X before the hexadecimal value to indicate its hexadecimal
	fmt.Printf("%o \n", 555555)
	fmt.Printf("%#o", 555555) // # will add a leading 0 before the ocatal value to indicate its ocatal

}

// 44 11100 4b56499 A736F507B
// 74f3
// 0x74f3
// 2075043
// 02075043
