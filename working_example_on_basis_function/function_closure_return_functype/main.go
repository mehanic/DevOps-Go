package main

import (
	"fmt"
)

func wrapper() func() int {
	x := 40
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper
	fmt.Println(increment())
	fmt.Println(increment())
}

// func wrapper return an anonymous function which inturn returns a type of int