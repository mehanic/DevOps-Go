package main

import "fmt"

func getSomeVars() string {
	fmt.Println("getSomeVars execution")
	return "getSomeVats result"
}

func main() {
	defer fmt.Println("After work")
	defer func() {
		fmt.Println(getSomeVars())
	}()
	fmt.Println("Some userful work")
}

// Some userful work
// getSomeVars execution
// getSomeVats result
// After work
