package main

import (
	"fmt"
)

func main() {
	hello("Shyam")
	s1 := foo("Shyam")
	fmt.Println(s1)
	x, y := mouse("Shyam", "Patel")
	fmt.Println(x)
	fmt.Println(y)
}

// everything in go is PASS BY VALUE
func hello(s string) {
	fmt.Println("Hello, ", s)
}

func foo(s string) string {
	return fmt.Sprint("Hello from foo, ", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ", says hello from mouse")
	b := false
	return a, b
}

// Hello,  Shyam
// Hello from foo, Shyam
// Shyam Patel, says hello from mouse
// false

