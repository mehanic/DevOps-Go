package main

import (
	"fmt"
	"reflect"
)

// Define types similar to classes in Python
type Complex struct{}

type FooBar struct{}

func main() {
	foobar := FooBar{}
	fmt.Println("Type of foobar:", reflect.TypeOf(foobar)) // Type of foobar

	// The type of FooBar itself (struct type)
	fmt.Println("Type of FooBar:", reflect.TypeOf(FooBar{}))

	// Type assertion in Go
	fmt.Println("foobar is of type FooBar:", reflect.TypeOf(foobar) == reflect.TypeOf(FooBar{}))

	// Reflection to check if a type implements an interface or is a struct
	fmt.Println("Is FooBar a struct:", reflect.TypeOf(FooBar{}).Kind() == reflect.Struct)
	fmt.Println("Is Complex a struct:", reflect.TypeOf(Complex{}).Kind() == reflect.Struct)
}

