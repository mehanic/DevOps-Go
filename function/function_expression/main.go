package main

import "fmt"

func main() {

	// the only way to declare a func inside a func
	// anonymous function
	greeting := func() {
		fmt.Println("greeting!!")
	}
	fmt.Printf("%T\n", greeting)

	greeting()

	fmt.Println(makeGreeter()())
}

// func makeGreeter returns a func that return a string
func makeGreeter() func() string {
	return func() string {
		return "Greeter made by maker"
	}
}
