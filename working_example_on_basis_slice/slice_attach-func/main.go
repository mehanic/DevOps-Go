package main

import "fmt"

type mySlice []int

func (s mySlice) print() {
	fmt.Println(s)
}

func main() {

	xi := []int{3, 5, 1, 8, 7, 9, 3, 0}

	// xi is converted to type mySlice, func print is attached to it
	mySlice(xi).print()
}
