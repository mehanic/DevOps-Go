package main

import (
	"fmt"
)

// sayHello returns a function that prints the sayHello's parameter
func sayHello(name string) func() {
	return func() {
		fmt.Printf("Hello %s", name)
	}
}
func main() {
	f := sayHello("Sharad")
	f() // Print : Hello Sharad
}
