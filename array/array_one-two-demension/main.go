package main

import "fmt"

func main() {

	arr := [5]int{}         // arr is an array if there has number in [] when defined, otherwise arr is a slice
	fmt.Println(arr)        // [0,0,0,0,0]
	fmt.Println(len(arr))   // 5
	fmt.Printf("%T\n", arr) // [5]int

	twoDimArr := [2][3]int{}      // multi-dimension
	fmt.Println(twoDimArr)        // [[0,0,0] [0,0,0]]
	fmt.Println(len(twoDimArr))   // 2
	fmt.Printf("%T\n", twoDimArr) // [2][3]int
}

// array is not dynamic, should be fixed size when defined an array
// array is less used in GO, usually use slice which is dynamic

// Arrays in Golang are value types unlike other languages like C, C++, and Java where arrays are reference types.
