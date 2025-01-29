package main

import "fmt"

func main() {

	var xs []string = []string{"a", "b", "c", "d", "e", "f"}
	var xs2 []string = []string{"g", "h", "i", "j", "k", "l"}

	fmt.Println("xs", xs)
	fmt.Println("xs2", xs2)

	fmt.Println(append(xs, "z")) // append returns a new slice without modifies the old one
	fmt.Println(xs)

	fmt.Println(append(xs, xs2...)) // append slices
	fmt.Println(append(xs[2:], xs2[:3]...))

}
