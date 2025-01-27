package main

import "fmt"

func main() {
	var speed1 int

	fmt.Println(speed1)

	// CORRECT DECLARATIONS
	var speed int
	var SpeeD int

	// underscore is allowed but not recommended
	var _speed int

	// Unicode letters are allowed
	var 速度 int

	// keep the compiler happy
	_, _, _, _ = speed, SpeeD, _speed, 速度

	println("---------------------")

	var nFiles int
	var counter int
	var nCPU int

	fmt.Println(
		nFiles,
		counter,
		nCPU,
	)

	println("---------------------")

	var heat float64
	var ratio float64
	var degree float64

	fmt.Println(
		heat,
		ratio,
		degree,
	)
	println("---------------------")

	var off bool
	var valid bool
	var closed bool

	fmt.Println(
		off,
		valid,
		closed,
	)
	println("---------------------")

	var msg string
	var name string
	var text string

	fmt.Println(
		msg,
		name,
		text,
	)

	println("---------------------")

	var speed2 int    // numeric type
	var heat2 float64 // numeric type
	var off2 bool
	var brand2 string

	fmt.Println(speed2)
	fmt.Println(heat2)
	fmt.Println(off2)

	println("---------------------")

	var (
		speed3 int
		heat3  float64

		off3   bool
		brand3 string
	)

	fmt.Println(speed3)
	fmt.Println(heat3)
	fmt.Println(off3)
	fmt.Printf("%q\n", brand3)

	fmt.Printf("%q\n", brand2)
}

// 0
// ---------------------
// 0 0 0
// ---------------------
// 0 0 0
// ---------------------
// false false false
// ---------------------

// ---------------------
// 0
// 0
// false
// ---------------------
// 0
// 0
// false
// ""
// ""
