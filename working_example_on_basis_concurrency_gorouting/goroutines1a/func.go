package goroutines

import (
	"fmt"
)

/*

A Goroutine is a function or method which executes independently and simultaneously in connection
with any other Goroutines present in our program. The cost of creating Goroutines is very small as
compared to the thread. Every program contains at least a single Goroutine and that
Goroutine is known as the main Goroutine. All the Goroutines are working under the main Goroutines
if the main Goroutine terminated, then all the goroutine present in the program also terminated.

*/

func Greetings(name string) {
	for i := 0; i < len(name); i++ {
		fmt.Println("Welcome ", name)
	}

}
