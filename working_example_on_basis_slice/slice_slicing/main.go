package main

import "fmt"

func main() {

	var xs []string = []string{"a", "b", "c", "d", "e", "f"}

	fmt.Println("xs", xs)
	fmt.Println("len(xs)", len(xs))

	fmt.Println("xs[1]", xs[1]) // index access

	fmt.Println("xs[1:3]", xs[1:3]) // slicing
	fmt.Println("xs[1:]", xs[1:])   // slicing

	fmt.Println("this is a string"[2]) // string is a slice of rune

}
