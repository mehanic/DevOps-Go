package main

import "fmt"

func main() {

	a := 3
	fmt.Println(a)  // value of var a
	fmt.Println(&a) // memory address of var a

	var b *int = &a // assign address of a to b, now b is a pointer to a, *int means type pointer to int
	fmt.Println(b)  // value of b, is the address of var a
	fmt.Println(&b) // address of b
	fmt.Println(*b) // the value saved in the address which b pointer to, value of var a, (*) is a dereference operater

	*b++            // change the value which b is pointer to, value of var a
	fmt.Println(a)  // 4
	fmt.Println(*b) // 4

	a++             // change value of a
	fmt.Println(a)  // 5
	fmt.Println(*b) // 5
}
