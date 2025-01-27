package main

import (
	"fmt"
)

var a string = "this is string"                                            // package scope; declaring and assigning at the same time
var b, c string = "this value is stored in b", "this value is stored in c" // package scope
var d string

func main() {
	fmt.Printf("%v \n", a) // %v will format apply the default formating type of the given type
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)

	d = "this value is stored in d"
	fmt.Printf("%v \n", d) // this variable is declared in package scope and assigned the value in main func

	var e int = 42 // variable declared inside func with type and intialized
	fmt.Printf("%v \n", e)

	f := 47 // variable declared inside func without type and intialized. Go detects type automatically
	fmt.Printf("%v \n", f)
	fmt.Printf("%T \n", f) // Prints the type of the variable

	g, h := "this is stored in g", "this is stored in h"
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)

	i, j, k, l := "this is stored in i", 23, true, 25.789
	fmt.Printf("%v \n", i)
	fmt.Printf("%v \n", j)
	fmt.Printf("%v \n", k)
	fmt.Printf("%v \n", l)

	m := "this is from m"
	fmt.Printf("%v \n", m)

	n := `this is a value stored in n`
	//The back quotes are used to create raw string literals which can contain any type of character
	//Raw string literals are character sequences between back quotes ``. Within the quotes, any character is legal except back quote.
	fmt.Printf("%v \n", n)
}

// this is string
// this value is stored in b
// this value is stored in c
// this value is stored in d
// 42
// 47
// int
// this is stored in g
// this is stored in h
// this is stored in i
// 23
// true
// 25.789
// this is from m
// this is a value stored in n
