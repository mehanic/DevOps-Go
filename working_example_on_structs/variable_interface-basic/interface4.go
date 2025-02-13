package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = "hello"
	fmt.Println("Type of x:", reflect.TypeOf(a))
	fmt.Println(a)
	s := a.(string)
	fmt.Println(s)

	s, ok := a.(string)
	fmt.Println(s, ok)
	f, ok := a.(float32)
	fmt.Println(f, ok)

	// g := a.(float64)
	// fmt.Println(g)

}

// Type of x: string
// hello
// hello
// hello true
// 0 false
