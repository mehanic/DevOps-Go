package main

import "fmt"

// in Go, we do not all caps const identifiers
const p = "death & taxes"

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}

// a CONSTANT is a simple unchanging value

// p -  death & taxes
// q -  42
