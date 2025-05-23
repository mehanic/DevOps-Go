package main

import "fmt"

func main() {

	fmt.Println()

	if true {
		fmt.Println("here ran")
	}

	if false {
		fmt.Println("there did not run")
	}

	fmt.Println()

	b := true
	if foo := "initialize"; b {
		fmt.Println(foo) // foo only exist here
	}
	// fmt.Println(foo) // foo is not exist here

	fmt.Println()

	c := 3
	if c == 1 {
		fmt.Println("if")
	} else if c == 2 {
		fmt.Println("else if")
	} else if c == 3 {
		fmt.Println("could have many else-if after if")
	} else {
		fmt.Println("else")
	}
	fmt.Println()

	if c == 1 {
		fmt.Println("if")
	} else {
		fmt.Println("could only one else after if")
	}
	fmt.Println()
}



// here ran

// initialize

// could have many else-if after if

// could only one else after if

