package main

import (
	"fmt"
	"reflect"
)

func main() {
	k := 4
	fmt.Println(reflect.TypeOf(k))
	l := float64(k)
	fmt.Println(reflect.TypeOf(l))
}

// int
// float64
