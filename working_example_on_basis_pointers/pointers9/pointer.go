package main

import "fmt"

func main() {
	var ptr *int // pointer with no value
	fmt.Println(ptr)

	var name string = "John"
	var namePtr *string = &name

	fmt.Println("name pointer is ", namePtr)

}
