package defers

import "fmt"

func SendPanic() {

	/*
		There are certain operations in Go that automatically return panics and stop the program.
		Common operations include indexing an array beyond its capacity, performing type assertions,
		calling methods on nil pointers, incorrectly using mutexes, and attempting to work with
		closed channels. Most of these situations result from mistakes made while programming that
		the compiler has no ability to detect while compiling your program.
	*/

	defer Test() // Test function only works if we defer it.

	panic("This is an error.")

	Test() // this won't work.

	defer fmt.Println("THIS WON'T WORK TOO!!")

}

func Test() {
	fmt.Println("This is a test.")
}
