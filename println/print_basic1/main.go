package main

import "fmt"

// lesson 002: an uint16 is a number from 0 to 65535

func main() {
	fmt.Println("")
	fmt.Println("an uint16 is a number from 0 to 65535")
	ui1 := uint16(0)
	ui2 := uint16(65535)
	fmt.Println(ui1, ui2)
	fmt.Println("")
	fmt.Println("If you go below 0 or above 65535 watch what happens")
	ui1 = ui1 - 1
	ui2 = ui2 + 1
	fmt.Println(ui1, ui2)
	fmt.Println("")
	fmt.Println("This is easier to understand if we look at the binary zeros and ones")
	fmt.Printf("%b %b\n", ui1, ui2)
	fmt.Println("")
	fmt.Println("We had eight positions from lesson 001 with a byte and now we have doubled that")
	fmt.Println("")
	fmt.Println("32768 16384 8192 4096 2048 1024 512 256 128 64 32 16 8 4 2 1")
	fmt.Println("    1     1    1    1    1    1   1   1   1  1  1  1 1 1 1 1")
	fmt.Println("")
	fmt.Println("32768+16384+8192+4096+2048+1024+512+256+128+64+32+16+8+4+2+1=65535")
	fmt.Println("")
	fmt.Println("The u in uint16 stands for unsigned, meaning no negative numbers.")
	fmt.Println("An int16 is a number from -32768 to 32767")
	fmt.Println("")
	i1 := int16(-32768)
	i2 := int16(32767)
	fmt.Println(i1, i2)
	fmt.Printf("%b %b\n", i1, i2)
	i1 = i1 - 1
	i2 = i2 + 1
	fmt.Println(i1, i2)
	fmt.Printf("%b %b\n", i1, i2)
	fmt.Println("")
	fmt.Println(" S 16384 8192 4096 2048 1024 512 256 128 64 32 16 8 4 2 1")
	fmt.Println(" 1     1    1    1    1    1   1   1   1  1  1  1 1 1 1 1")
	fmt.Println("")
	fmt.Println("16384+8192+4096+2048+1024+512+256+128+64+32+16+8+4+2+1=32767")
	fmt.Println("")
	fmt.Println("When dealing with negative numbers we lose that 32768 position.")
	fmt.Println("We need that position to tell us positive or negative.")
	fmt.Println("")

	// uncomment the code below to see a loop run

	/*
		   	fmt.Println("")
				i := int16(-32768)
				for {
					fmt.Printf("%b %d\n", i, i)
					i += 1000
					if i > 1000 {
						break
					}
				}
	*/
}


// an uint16 is a number from 0 to 65535
// 0 65535

// If you go below 0 or above 65535 watch what happens
// 65535 0

// This is easier to understand if we look at the binary zeros and ones
// 1111111111111111 0

// We had eight positions from lesson 001 with a byte and now we have doubled that

// 32768 16384 8192 4096 2048 1024 512 256 128 64 32 16 8 4 2 1
//     1     1    1    1    1    1   1   1   1  1  1  1 1 1 1 1

// 32768+16384+8192+4096+2048+1024+512+256+128+64+32+16+8+4+2+1=65535

// The u in uint16 stands for unsigned, meaning no negative numbers.
// An int16 is a number from -32768 to 32767

// -32768 32767
// -1000000000000000 111111111111111
// 32767 -32768
// 111111111111111 -1000000000000000

//  S 16384 8192 4096 2048 1024 512 256 128 64 32 16 8 4 2 1
//  1     1    1    1    1    1   1   1   1  1  1  1 1 1 1 1

// 16384+8192+4096+2048+1024+512+256+128+64+32+16+8+4+2+1=32767

// When dealing with negative numbers we lose that 32768 position.
// We need that position to tell us positive or negative.

