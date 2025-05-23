package main

import (
	"fmt"
)

func main1() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // deferencing

	*b = 45

	fmt.Println(a)
}
