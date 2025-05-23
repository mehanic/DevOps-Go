package main

import "fmt"

func PrintVariables() {

	// 1. Unsigned Integers

	/*
		NERD FACT
		 uint8 - fit into 8 bits of memory -> 255
		 uint16 - fit into 16 bits of memory -> 65,535
		 uint32 -fit into 32 bits of memory -> 4,294,967,295
		 uint64 - fit into 64 bits of memory -> 18,446,744,073,709,551,615
		 uint - unsigned. Depends on which system are your running, 32-bit or 64-bit system.

		 In order not to waste memory, it is advised to assign to use uint8 unless your program required it
	*/

	// uint8, uint16, uint32, uint64, uint

	var roomNumber uint8 = 10
	//will throw error
	// var roomNumber2 uint8 = 256

	fmt.Printf("UINT8, %d\n", roomNumber)

	// 2. Integers

	// both positive and negative integer.
	var roomNumber3 int8 = 127
	fmt.Printf("INT8, %d\n", roomNumber3)
	// will throw error
	//var roomNumber4 int8 = 128

	// 3. Float

	const longitude = 24.806078

	// 0.2f used to print 2 significant figures
	fmt.Printf("FLOAT, %0.2f\n", longitude)

	// 4. String, runes

	var text string = "Hello"
	fmt.Printf("String, %v\n", text)

	// 5. boolean values

	var isTrue bool = true
	isFalse := false
	fmt.Printf("bool, %v, %v\n", isTrue, isFalse)

}

func main() {
	PrintVariables()
}

// UINT8, 10
// INT8, 127
// FLOAT, 24.81
// String, Hello
// bool, true, false
