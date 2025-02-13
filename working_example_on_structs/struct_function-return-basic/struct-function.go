package main

import "fmt"

const o string = "name"

type s struct {
	c string
	b string
}

func main() {
	// var t string     //init
	// var f = "sadsad" //zapisat
	// z := "asdad"     //init i zapis
	a, b := test()
	fmt.Println(a, b)
}

func test() (string, string) {
	a := "hello"
	b := "world"
	return a, b
}

// hello world
