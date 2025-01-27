package main

import "fmt"

//> < <= >= ==
// and &&
// or ||

func main() {
	a := 3
	b := 5
	c := 4
	d := 5
	var k bool
	k = a > b || c == d
	fmt.Println(k)
}

// false
