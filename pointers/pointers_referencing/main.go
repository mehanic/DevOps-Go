package main

import (
	"fmt"
)

func main() {

	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	//
	//variable b is a pointer to memory address (*int) where the int is stored
}
