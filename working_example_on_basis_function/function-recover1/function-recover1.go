package main

import "fmt"

func foo() {
	defer fmt.Println("Bye foo")
	fmt.Println("Hello foo")
	bar()
	fmt.Println("After bar")
}

func bar() {
	defer fmt.Println("Bye bar")
	fmt.Println("Hello bar")
	baz()
	fmt.Println("After baz")
}

func baz() {
	defer fmt.Println("Bye baz")
	fmt.Println("Hello baz")
	panic("panic in baz")
}

func main() {
	foo()
	fmt.Println("After foo")
}
