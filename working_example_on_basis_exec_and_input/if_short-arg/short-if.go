package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("42")

	if err == nil {
		fmt.Println("There was no error, n is", n)
	}
}

func main1() {
	if n, err := strconv.Atoi("42"); err == nil {
		// n and err are available here
		fmt.Println("There was no error, n is", n)
	}

	// n and err are not available here
	// fmt.Println(n, err)
}

func main2() {
	if a := os.Args; len(a) != 2 {

		// only a is available here
		fmt.Println("Give me a number.")

		// no need to return on error
		// return

	} else if n, err := strconv.Atoi(a[1]); err != nil {

		// a, n and err are available here
		fmt.Printf("Cannot convert %q.\n", a[1])

		// no need to return on error
		// return

	} else {
		// all of the variables (names)
		// are available here
		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
	}

	// a, n and err are not available here
	// they belong to the if statement

	// TRY:
	// fmt.Println(a, n, err)
}

func main3() {
	// UNCOMMENT THIS TO SEE IT IN ACTION:
	// var n int

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		// else if here shadows the main func's n
		// variable and it declares it own
		// with the same name: n

		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
	}

	// n here belongs to the main func
	// not to the if statement above

	// UNCOMMENT ALSO LINES BELOW TO SEE IT IN ACTION:
	// fmt.Printf("n is %d. 👻 👻 👻 - you've been shadowed ;-)\n", n)
}

func main4() {
	// n and err belongs to the main function now
	// not only to the if statement below
	var (
		n   int
		err error
	)

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number.")
	} else if n, err = strconv.Atoi(a[1]); err != nil {
		// here else if uses the main func's n and err
		// variables instead of its own

		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		// assigns the result back into main func's
		// n variable
		n *= 2

		fmt.Printf("%s * 2 is %d\n", a[1], n)
	}

	// if statement has calculated the result using
	// the main func's n variable
	//
	// so, that's why it prints the correct result here
	fmt.Println("n is", n)
}


//6 * 2 is 12
//n is 12
