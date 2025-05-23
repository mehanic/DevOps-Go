package main

import (
	"fmt"
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	//
	// each int element is 4 bytes (on 64-bit)
	//
	// let's say the ages point to 1000th.
	// ages[1:] will point to 1004th
	// ages[2:] will point to 1008th and so on.
	//
	// they all will be looking at the same
	// backing array.
	//

	ages := []int{35, 15, 25}
	red, green := ages[0:1], ages[1:3]

	s.Show("ages", ages)
	s.Show("red", red)
	s.Show("green", green)

	fmt.Println(red[0])
	// fmt.Println(red[1]) // error
	// fmt.Println(red[2]) // error

	{
		var ages []int
		s.Show("nil slice", ages)

		// or just:
		s.Show("nil slice", []int(nil))
	}
}

//--------

// type collection [4]string // #1
type collection []string // #2

// go is pass by copy
// only the slice header is copied: 3 integer fields (24 bytes)
// think of passing an array with millions of elements instead.

func main1() {
	// SliceHeader lives here:
	// https://golang.org/src/runtime/slice.go

	s.PrintElementAddr = true

	// #1
	data := collection{"slices", "are", "awesome", "period", "!!" /* #5 */}

	// data := collection{"slices", "are", "awesome", "period", "!!"}

	change(data) // #1

	s.Show("main's data", data)                           // #1
	fmt.Printf("main's data slice's header: %p\n", &data) // #3

	// ----------------------------------------------------------------------
	// #4
	array := [...]string{"slices", "are", "awesome", "period", "!!" /* #5 */}

	// array's size depends on its elements
	fmt.Printf("array's size: %d bytes.\n", unsafe.Sizeof(array))

	// slice's size is always fixed: 24 bytes (on a 64-bit system) — slice value = slice header
	fmt.Printf("slice's size: %d bytes.\n", unsafe.Sizeof(data))
}

// #1
// passed value will be copied in the function
func change(data collection) {
	// data is a new variable inside the function:
	// var data collection

	data[2] = "brilliant!"

	s.Show("change's data", data)
	fmt.Printf("change's data slice's header: %p\n", &data) // #3
}
