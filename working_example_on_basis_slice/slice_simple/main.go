package main

import "fmt"

func main() {

	xi := []int{} // xi is defined without number in [] is a slice

	fmt.Println(xi)
	fmt.Println(len(xi))

	xi = append(xi, 1) // append 1 at last of xi and return a new slice

	fmt.Println(xi)
	fmt.Println(len(xi))

}