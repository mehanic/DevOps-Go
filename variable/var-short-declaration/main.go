package main

import "fmt"

func main() {
	// := allows you to declare and initialize at once
	a := 10
	b := "golang"
	c := 4.17
	d := true

	// %v: the value in a default format
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}

// declare, assign, initialize
// declare: var b string (declare b as a string)
// assign: b = ""
// initialize: b := "" (declare and assign at once)
