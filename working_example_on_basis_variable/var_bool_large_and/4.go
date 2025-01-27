package main

import "fmt"

//> < <= >= ==
// and &&
// or ||

func main() {
	a := 6
	b := 5
	c := 5
	d := 5
	var k bool
	k = a > b && c == d
	fmt.Println(k)
}

// true
