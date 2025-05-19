package defers

import (
	"fmt"
	"log"
)

/*

Go makes it possible to recover from a panic, by using the recover built-in function.
A recover can stop a panic from aborting the program and let it continue with execution instead.

An example of where this can be useful: a server wouldn’t want to crash if one of the client
connections exhibits a critical error. Instead, the server would want to close that connection
and continue serving other clients. In fact, this is what Go’s net/http does by default for HTTP servers.
*/

func RecoverPanic() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("Panic recovered\n", err)
		}
	}()

	defer fmt.Println(divideByZero(5, 2))
	fmt.Println(divideByZero(5, 0))

}

func divideByZero(a, b int) int {

	if b == 0 {
		panic("Zero Division error")
	}
	return a / b
}
