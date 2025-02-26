package main

import "fmt"

func main() {

	defer hello() // defer will run the func last
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2") // defer func will run in LIFO order

	world()

	if true {
		return // defer func will run before return
	}
	world()

	// output:
	// world
	// 2
	// 1
	// hello
}

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}
