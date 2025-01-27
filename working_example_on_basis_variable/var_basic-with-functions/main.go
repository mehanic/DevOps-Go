package main

import (
	"fmt"
)

func main() {
	// Variables
	variables()

	// Types & Zero Values
	typesAndZeroValues()

	// Naming Conventions
	namingConventions()

	// Go fmt Package
	// https://golang.org/pkg/fmt/
	// Printf
	printFExamples()

	constants()
	untypedConstants()

	dataTypes()
	compositeDataTypes()

	// Pointers
	pointers()
}

func variables() {
	var a int = 10
	var b float32 = 3.14
	const c string = "Hello World"

	var (
		d = 10
		e = 7
	)

	x, y := 10, 20

	fmt.Println(a, b, c, d, e, x, y)

	// In Go you cannot declare a variable and not use it.
	var name = "Edward"
	_ = name

	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	var (
		salary    float64 = 10000.0
		firstName string  = "Edward"
		gender    bool    = true
	)

	// salary = 10000.0
	// firstName = "Edward"
	// gender = true

	fmt.Println(salary, firstName, gender)

	// Variable Swapping
	var firstVar, SecondVar int = 5, 10
	firstVar, SecondVar = SecondVar, firstVar
	fmt.Println(firstVar, SecondVar)
}

func typesAndZeroValues() {
	var a int = 4
	var b float64 = 5.2

	fmt.Println("Types & Zero Vals", a, b)

	a = int(b)
	fmt.Println("Types & Zero Vals", a, b)
}

func namingConventions() {
	// It is better to use short versions of variable names.
	var mv int = 10        // mv = max value
	var max_value int = 10 // not good practice.
	_ = max_value

	var packetsReceived int // not good practice.
	_ = packetsReceived
	var pr int // pr = packets received

	const MAX_VALUE = 100 // not good practice.
	_ = MAX_VALUE
	const maxVal = 100 // maxVal = max value

	fmt.Println(mv, pr, maxVal)

	// If you declare a variable starting with an uppercase letter, it will be exported.
	var Max = 100
	fmt.Println(Max)

	writeToDB := true  // ok
	writeToDb := false // not ok

	fmt.Println(writeToDB, writeToDb)
}

func printFExamples() {
	const pi float64 = 3.14159265358979323846
	x := 5
	isbool := true
	figure := "circle"

	fmt.Printf("Value of pi: %f \n", pi)           // %f = float
	fmt.Printf("Value of pi: %T \n", pi)           // %t = type
	fmt.Printf("Value of x: %d \n", x)             // %d = decimal
	fmt.Printf("Value of x: %+d \n", x)            // %+d = decimal with sign
	fmt.Printf("Value of isbool: %t \n", isbool)   // %t = boolean
	fmt.Printf("Value of isbool: %T \n", isbool)   // %T = type
	fmt.Printf("Value of pi: %.7f \n", pi)         // %.2f = float with 2 decimal places
	fmt.Printf("Value of pi: %9.2f \n", pi)        // %9.2f = float with 2 decimal places and 9 characters (spaces) wide
	fmt.Printf("Value of pi: %9.f \n", pi)         // %9.f = float with 9 characters (spaces) wide
	fmt.Printf("Value of pi: %s \n", figure)       // %s = string
	fmt.Printf("Value of pi: %q \n", figure)       // %q = quoted string
	fmt.Printf("Value of pi: %v \n", figure)       // %v = value (any type)
	fmt.Printf("Value of int in Base 2: %b \n", x) // %b = base 2
	fmt.Printf("Value of bit: %x \n", 100)

	var a, y int
	fmt.Printf("%T -> %d\n", a, y)
}

func constants() {
	// Constants must always be initialized.
	const days int = 7
	_ = days

	const pi float64 = 3.14
	_ = pi

	const secondsInHour = 3600

	duration := 234 // in hours

	fmt.Printf("Duration in seconds: %d\n", duration*secondsInHour)

	//x, y = 5, 0
	//fmt.Println(x / y) // panic: runtime error: integer divide by zero

	//const a, b int = 5, 0
	//fmt.Println(a / b) // panic: runtime error: integer divide by zero

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)

	fmt.Println(min1, min2, min3)

	// Constants must be initialized before runtime.
	// const power = math.Pow(2, 3)
	// fmt.Println(power)

	// Variables cannot be used to initialize constants.

	// You cannot change a constant and you cannot add an array to a constant because it is not a constant expression and its values can change.
	//const y = [2]int{5, 6}
}

func untypedConstants() {
	const a float64 = 5.1 // Typed constant - type float64

	const b = 6.7 // Untyped constant - type float64

	const c float64 = a * b // Typed constant - type float64 (5.1 * 6.7)
	fmt.Printf("This is a %f\n", c)

	const d = 5 > 10 // Untyped constant - this will be a logical check (false)
	fmt.Printf("d is %t\n", d)

	// Using untyped constants gives you more flexibility.
	const e = 10
	fmt.Printf("e is %v and is type %T\n", e, e)

	var f int = e           // f is an int
	var g float64 = e * 2.2 // g is float64

	fmt.Printf("f is %v and is type %T\ng is %.2f and is type %T\n", f, f, g, g)

	// IOTA - special predeclared identifier - it auto increments a value by 1 or whatever value you set it to.
	const (
		c1 = iota * 2 // 0
		c2
		c3
	)

	fmt.Println(c1, c2, c3)

	const (
		x = -(iota + 2) // -2
		_               // -3 - if you want to skip a value, use the blank identifier.
		y               // -4
		z               // -5
	)

	fmt.Println(x, y, z)
}

func dataTypes() {
	var i1 int8 = 100
	fmt.Printf("%T\n", i1)

	var i2 int16 = 10000
	fmt.Printf("%T\n", i2)

	var i3 int32 = 1000000000
	fmt.Printf("%T\n", i3)

	var i4 int64 = 1000000000000000000
	fmt.Printf("%T\n", i4)

	var i5 int = 1000000000000000000
	fmt.Printf("%T\n", i5)

	var int8 int8 = -128 // Increasing this value by 1 will cause an overflow.
	fmt.Printf("%T\n", int8)

	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	var b1 byte = 2
	var b2 byte = 255 // Increasing this value by 1 will cause an overflow.
	fmt.Printf("%T %T\n", b1, b2)

	var r rune = 'f'
	fmt.Printf("%v\n", r)

	var b bool = true
	fmt.Printf("%T\n", b)

	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)
}

func compositeDataTypes() {
	var s [2]string
	s[0] = "Hello"
	s[1] = "Go!"
	fmt.Printf("%v\n", s)

	var numbers = [4]int{1, 2, 3, 4}
	fmt.Printf("%v\n", numbers)

	var st = []int{1, 2, 3, 4}
	fmt.Printf("%T, %v\n", st, st)

	var mt = map[string]int{"foo": 42}
	fmt.Printf("%T, %v\n", mt, mt)

	var mt1 = map[int]int{0: 1000}
	fmt.Printf("%T, %v\n", mt1, mt1)

	// Structs are used to create custom data types.
	type Person struct {
		name string
		age  int
	}

	var p = Person{name: "Edward", age: 35}
	fmt.Printf("%T, %v\n", p, p)
}

func pointers() {
	var anInt int = 42
	var p1 *int = &anInt
	fmt.Printf("anInt is type %T and its value is %v\n", anInt, anInt)
	fmt.Printf("anInt is type %T and its value is %v\n", &anInt, &anInt)
	fmt.Printf("p1 is type %T and its address value is %v\n", p1, p1)
	fmt.Printf("p1 is type %T and its value is %v\n", *p1, *p1)

	p2 := 5
	fmt.Println(p2, &p2)

	changePointerValue(&p2)
	fmt.Println(p2)

	fmt.Printf("The type of functionType is %T\n", functionType)
}

func functionType() {
	// Intentionally left blank for demo of function type.
}

func changePointerValue(x *int) {
	*x = 10
}

// 10 3.14 Hello World 10 7 10 20
// Audi 50000
// 10000 Edward true
// 10 5
// Types & Zero Vals 4 5.2
// Types & Zero Vals 5 5.2
// 10 0 100
// 100
// true false
// Value of pi: 3.141593
// Value of pi: float64
// Value of x: 5
// Value of x: +5
// Value of isbool: true
// Value of isbool: bool
// Value of pi: 3.1415927
// Value of pi:      3.14
// Value of pi:         3
// Value of pi: circle
// Value of pi: "circle"
// Value of pi: circle
// Value of int in Base 2: 101
// Value of bit: 64
// int -> 0
// Duration in seconds: 842400
// -500 -300 100
// This is a 34.170000
// d is false
// e is 10 and is type int
// f is 10 and is type int
// g is 22.00 and is type float64
// 0 2 4
// -2 -4 -5
// int8
// int16
// int32
// int64
// int
// int8
// float64 float64 float64
// uint8 uint8
// 102
// bool
// string
// [Hello Go!]
// [1 2 3 4]
// []int, [1 2 3 4]
// map[string]int, map[foo:42]
// map[int]int, map[0:1000]
// main.Person, {Edward 35}
// anInt is type int and its value is 42
// anInt is type *int and its value is 0xc000012188
// p1 is type *int and its address value is 0xc000012188
// p1 is type int and its value is 42
// 5 0xc000012190
// 10
// The type of functionType is func()
