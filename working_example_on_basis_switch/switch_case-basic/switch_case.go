package main

import "fmt"

func main() {
	a := 0
	fmt.Printf("Hello %d", a)

	for a < 10 {

		i := isTest(a)
		switch i {
		case 1:
			fmt.Println("hello")
		default:
			fmt.Println("hello")
		}

		switch {
		case i == 2:
			fmt.Println("hello")
		default:
			fmt.Println("hello")
		}

		if i := isTest(a); i == 1 {

		}
		if a == 50 {
			fmt.Println("hello")
		} else {
			fmt.Println("world")
		}
		a++
	}

}

func isTest(a int) int {
	if a == 2 {
		return 1
	} else if a == 3 {
		return 3
	}

	return 4
}


// Hello 0hello
// hello
// world
// hello
// hello
// world
// hello
// hello
// world
// hello
// hello
// world
// hello
// hello
// world
// hello
// hello
// world
// hello
// hello
// world
// hello
// hello
// world
// hello
// hello
// world
// hello
// hello
// world

