package main

import (
	"fmt"
)

func main() {
	var usrStr string
	fmt.Print("Enter your string: ")
	fmt.Scanln(&usrStr)

	index := 0
	for _, eachChar := range usrStr {
		fmt.Printf("%c --> %d\n", eachChar, index)
		index = index + 1
	}
}


// Enter your string: first
// f --> 0
// i --> 1
// r --> 2
// s --> 3
// t --> 4
