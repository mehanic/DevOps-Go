package main

import (
	"fmt"
)

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2
}

type InterfaceHere interface {
	Sum() int
}

func main() {
	var a InterfaceHere
	sh := structHere{1, 2}
	os := otherStruct{2, 2}
	a = &sh
	fmt.Println(a.Sum())
	a = os
	fmt.Println(a.Sum())

	var i InterfaceHere = otherStruct{2, 2}
	i.Sum()
}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}

// 3
// 4
