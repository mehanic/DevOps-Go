package main

import "fmt"

func main() {
	a := 300

	ptr := &a

	fmt.Println(a)
	fmt.Println(*ptr)

	*ptr = 20

	fmt.Println(a)
	fmt.Println(*ptr)
}
