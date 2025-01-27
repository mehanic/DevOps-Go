package main

import "fmt"

func main() {

	// ! (not)
	if !false {
		fmt.Println("this ran")
	}

	// || (or)
	if false || true {
		fmt.Println("this ran")
	}

	// && (and)
	if false && true {
		fmt.Println("this did not run")
	}
}
