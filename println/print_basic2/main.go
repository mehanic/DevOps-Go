package main

import "fmt"

// lesson 003: know your ints and appends

func main() {
	fmt.Println("")
	fmt.Println("byte   : 0 255")
	fmt.Println("uint16 : 0 to 65535")
	fmt.Println("uint32 : 0 to 4294967295")
	fmt.Println("uint64 : 0 to 18446744073709551615")
	fmt.Println("int16  : -32768 to 32767")
	fmt.Println("int32  : -2147483648 to 2147483647")
	fmt.Println("int64  : -9223372036854775808 to 9223372036854775807")
	fmt.Println("")
	fmt.Println("When you just say int it will default to either")
	fmt.Println("")
	fmt.Println("  int32 or int64")
	fmt.Println("")
	fmt.Println("Depending on the machine you are on. This is known as Platform dependent")

	test := uint64(18446744073709551615) // try chaning this number
	fmt.Printf("%b\n", test)

	fmt.Println("")
	fmt.Println("When deciding which int to use, think memory use.")
	fmt.Println("You only have so much ram. Why store a list of integers each with 64 bits?")
	fmt.Println("If that list is very long and you only need 8 bits for each number...")
	fmt.Println("")
	fmt.Println("Space for Zeros and Ones is precious!")
	fmt.Println("")
	fmt.Println("Now lets talk about lists or arrays.")
	fmt.Println("")

	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(list1)
	fmt.Println("")
	fmt.Println("I can add one item to the end")
	list1 = append(list1, 11)
	fmt.Println("")
	fmt.Println(list1)
	fmt.Println("")
	fmt.Println("I can add two items to the end")
	list1 = append(list1, 12, 13)
	fmt.Println("")
	fmt.Println(list1)
	fmt.Println("")
	list2 := []int{100, 200, 300}
	fmt.Println(list2)
	fmt.Println("")
	fmt.Println("I can add another list to end")
	list1 = append(list1, list2...)
	fmt.Println("")
	fmt.Println(list1)
	fmt.Println("")
	fmt.Println("I can add something to front of list")
	list1 = append([]int{0}, list1...)
	fmt.Println("")
	fmt.Println(list1)
	fmt.Println("")

}


// byte   : 0 255
// uint16 : 0 to 65535
// uint32 : 0 to 4294967295
// uint64 : 0 to 18446744073709551615
// int16  : -32768 to 32767
// int32  : -2147483648 to 2147483647
// int64  : -9223372036854775808 to 9223372036854775807

// When you just say int it will default to either

//   int32 or int64

// Depending on the machine you are on. This is known as Platform dependent
// 1111111111111111111111111111111111111111111111111111111111111111

// When deciding which int to use, think memory use.
// You only have so much ram. Why store a list of integers each with 64 bits?
// If that list is very long and you only need 8 bits for each number...

// Space for Zeros and Ones is precious!

// Now lets talk about lists or arrays.

// [1 2 3 4 5 6 7 8 9 10]

// I can add one item to the end

// [1 2 3 4 5 6 7 8 9 10 11]

// I can add two items to the end

// [1 2 3 4 5 6 7 8 9 10 11 12 13]

// [100 200 300]

// I can add another list to end

// [1 2 3 4 5 6 7 8 9 10 11 12 13 100 200 300]

// I can add something to front of list

// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 100 200 300]

