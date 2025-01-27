package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a float64
	a = 15
	b := 4.0
	c := a + b
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(b))
}

// 19
// float64
