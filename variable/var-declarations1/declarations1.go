package main

import "fmt"

var safe22 = true

func main() {
	fmt.Println(safe22)
	var speed, velocity int
	fmt.Println(speed, velocity)

	println("--------------")
	var myAge, yourAge int
	fmt.Println(myAge, yourAge)

	var temperature float64
	fmt.Println(temperature)

	var success bool
	fmt.Println(success)

	var language string
	fmt.Println(language)

	println("--------------")

	var height int
	fmt.Println(height)

	println("--------------")
	var isOn bool
	fmt.Println(isOn)
	println("--------------")

	var s string
	fmt.Printf("s (%T): %q\n", s, s)
	println("--------------")

	// integer types
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	// float types
	var f32 float32
	var f64 float64

	// complex types
	var c64 complex64
	var c128 complex128

	// bool type
	var b bool

	// string types
	var k string
	var r rune  // also a numeric type
	var by byte // also a numeric type

	fmt.Println(
		i, i8, i16, i32, i64,
		f32, f64,
		c64, c128,
		b, r, by,
	)

	// You could do it with Println as well
	fmt.Printf("%q\n", k)
	println("--------------")
	var (
		active bool
		delta  int
	)
	fmt.Println(active, delta)

	println("--------------")

	var firstName, lastName string = "", ""
	fmt.Printf("%q %q\n", firstName, lastName)
	println("--------------")
	//unused
	var isLiquid bool
	_ = isLiquid
	println("--------------")
	//short declaration
	var safe = true
	fmt.Println(safe)

	safe1 := true
	fmt.Println(safe1)

	println("--------------")

	// var name string = "Carl"
	// var name = "Carl"
	name := "Carl"

	// var isScientist bool = true
	// var isScientist = true
	isScientist := true

	// var age int = 62
	// var age = 62
	age := 62

	// var degree float64 = 5.
	// var degree = 5.
	degree := 5.

	fmt.Println(name, isScientist, age, degree)
	name2 := name
	fmt.Println(name2)
	println("--------------")

	println("--------------")

}

// true
// 0 0
// --------------
// 0 0
// 0
// false

// --------------
// 0
// --------------
// false
// --------------
// s (string): ""
// --------------
// 0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0
// ""
// --------------
// false 0
// --------------
// "" ""
// --------------
// --------------
// true
// true
// --------------
// Carl true 62 5
// Carl
// --------------
// -------------
