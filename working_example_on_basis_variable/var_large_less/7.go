package main

import "fmt"

func main() {
	a := 3
	b := 4
	if a > b {
		fmt.Println("a>b")
	} else if a == b {
		fmt.Println("a == b")
	} else if a < b {
		fmt.Println("a < b")
	}
}

// a < b
