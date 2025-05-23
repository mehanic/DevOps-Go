package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	fmt.Println(&a)
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
}

// 0xc0000ac010
// 0
// int
