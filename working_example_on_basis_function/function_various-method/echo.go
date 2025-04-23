
package main

import (
	"fmt"
	"os"
)

//
// Various method using Go
//
func main() {
	//data type declaration is different from Java, data types come after variable
	//When variable is not initialized implicitly, Go will automatically initialize them to the
	//default zero value
	var s string
	var sep = "."

	fmt.Println(len(os.Args))

	//
	// := short variable declaration, assigned type based on initializer values
	// ++ increment statement ,postfix
	// for is the only loop format in Go
	//
	// for initialization; condition ; post
	// initialization : simple statement, increment or assignment statement or a function call
	//
	// Short variable declaration
	//
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)

	//increment/decrement
	var i int = len(os.Args)
	for i--; i > 0; i-- {
		s += sep + os.Args[i]
	}
	fmt.Println(s)

	//function call
	for concate(); 1 > 2; {
	}

}

// print the arguments
func concate() {
	var s string
	var sep = "!"
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
}
