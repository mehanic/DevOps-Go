package main

import "fmt"

func main() {
	//variable (value,address)
	//
	fmt.Println("variable b")
	b := 3
	fmt.Println("value", b)
	fmt.Println("address ", &b)
	//pointer (*value, &address, link_to_another_address)
	var c *int
	c = &b
	fmt.Println("variable c")
	fmt.Println("value", *c)
	fmt.Println("address ", c)
	b = 5
	fmt.Println("variable b")
	fmt.Println("value", b)
	fmt.Println("address ", &b)
	fmt.Println("variable c after change")
	fmt.Println("value", *c)
	fmt.Println("address ", c)
}

// variable b
// value 3
// address  0xc00011a050
// variable c
// value 3
// address  0xc00011a050
// variable b
// value 5
// address  0xc00011a050
// variable c after change
// value 5
// address  0xc00011a050
