package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a) // Declares 'a' as a string and initializes it with "initial"

	var b, c int = 1, 2
	fmt.Println(b, c) // Declares 'b' and 'c' as integers with values 1 and 2

	var d = true
	fmt.Println(d) // Declares 'd' as a boolean with the value 'true'

	var e int
	fmt.Println(e) // Declares 'e' as an integer with no initial value, so it defaults to 0

	f := "apple"
	fmt.Println(f) // Uses shorthand notation to declare 'f' as a string and initialize it with "apple"
}

// initial
// 1 2
// true
// 0
// apple
