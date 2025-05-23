package main

import "fmt"

// lesson 001: a byte is a number from 0 to 255

func main() {
	fmt.Println("")
	fmt.Println("a byte is a number from 0 to 255")
	b1 := byte(0)
	b2 := byte(255)
	fmt.Println(b1, b2)
	fmt.Println("")
	fmt.Println("If you go below 0 or above 255 watch what happens")
	b1 = b1 - 1
	b2 = b2 + 1
	fmt.Println(b1, b2)
	fmt.Println("")
	fmt.Println("This is easier to understand if we look at the binary zeros and ones")
	fmt.Printf("%b %b\n", b1, b2) // with Printf vs Println and %b you can tell go to print the zeros and ones
	fmt.Println("")
	fmt.Println("A byte has 8 possible zeros or ones and 11111111 is the max")
	fmt.Println("If you add 1 it goes back to 00000000 or 0")
	fmt.Println("And if you are at 0 and subtract 1, it goes to 11111111")
	fmt.Println("")
	fmt.Println("128 64 32 16  8  4  2  1")
	fmt.Println("  1  1  1  1  1  1  1  1")
	fmt.Println("")
	fmt.Println("128+64+32+16+8+4+2+1=255")
	fmt.Println("")
	fmt.Println("Representing any number is just turning on or off each position")
	fmt.Println("")
	fmt.Println("128 64 32 16  8  4  2  1")
	fmt.Println("  1  0  0  0  0  0  0  1")
	fmt.Println("")
	fmt.Println("128+1=129")
	fmt.Println("")
	fmt.Println("Notice how 129 can be made with 128 on and everything else off except the last 1")
	fmt.Println("")
	b3 := byte(129)
	fmt.Println(b3) // change this to Printf and use %b to show 129 in binary
	fmt.Println("")
}


// a byte is a number from 0 to 255
// 0 255

// If you go below 0 or above 255 watch what happens
// 255 0

// This is easier to understand if we look at the binary zeros and ones
// 11111111 0

// A byte has 8 possible zeros or ones and 11111111 is the max
// If you add 1 it goes back to 00000000 or 0
// And if you are at 0 and subtract 1, it goes to 11111111

// 128 64 32 16  8  4  2  1
//   1  1  1  1  1  1  1  1

// 128+64+32+16+8+4+2+1=255

// Representing any number is just turning on or off each position

// 128 64 32 16  8  4  2  1
//   1  0  0  0  0  0  0  1

// 128+1=129

// Notice how 129 can be made with 128 on and everything else off except the last 1

// 129
