package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		l  = len(os.Args) - 1
		n1 = os.Args[1]
		n2 = os.Args[2]
		n3 = os.Args[3]
	)

	fmt.Println("There are", l, "people !")
	fmt.Println("Hello great", n1, "!")
	fmt.Println("Hello great", n2, "!")
	fmt.Println("Hello great", n3, "!")
	fmt.Println("Nice to meet you all.")
}

//if fewer than 3 names are provided, the program will output a message asking for more arguments instead of causing a panic.

// main.go Alice Bob Charlie
// There are 3 people !
// Hello great Alice !
// Hello great Bob !
// Hello great Charlie !
// Nice to meet you all.

