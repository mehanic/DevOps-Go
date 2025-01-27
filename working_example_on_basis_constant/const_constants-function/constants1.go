package main

import (
	"fmt"
	"time"
)

func main() {
	const min int = 1
	const pi float64 = 3.14
	const version string = "2.0.1"
	const debug bool = true
	const A rune = 'A'

	fmt.Println(min, pi, version, debug, A)

	main1()
	fmt.Println("------------")
	main3()
	fmt.Println("------------")

	main4()
	fmt.Println("------------")

	main5()
	fmt.Println("------------")

	main6()
	fmt.Println("------------")

	main7()
	fmt.Println("------------")

	main8()
	fmt.Println("------------")

	main9()
	fmt.Println("------------")

	main10()
	fmt.Println("------------")

	main11()
	fmt.Println("------------")

	main12()
	fmt.Println("------------")

	main13()
	fmt.Println("------------")

	main14()
	fmt.Println("------------")

	main15()
	fmt.Println("------------")

	main16()
	fmt.Println("------------")

	main17()
	fmt.Println("------------")

	main18()
}

func main1() {
	const min = 1
	const pi = 3.14
	const version = "2.0.1"
	const debug = true
	const A = 'A'

	fmt.Println(min, pi, version, debug, A)
}

func main3() {
	const min = 1 + 1
	const pi = 3.14 * min
	const version = "2.0.1" + "-beta"
	const debug = !true
	const A rune = 'A' + 1

	fmt.Println(min, pi, version, debug, A)
}

func main4() {
	// below declaration is the same as this one:

	// const (
	// 	min int = 1
	// 	max int = 1000
	// )

	const min, max int = 1, 1000

	fmt.Println(min, max)

	// print the types of min and max
	fmt.Printf("%T %T\n", min, max)
}

func main5() {
	// constants repeat the previous type and expression
	const (
		min int = 1
		max     // int = 1
	)

	fmt.Println(min, max)

	// print the types of min and max
	fmt.Printf("%T %T\n", min, max)
}

func main6() {
	// constants repeat the previous type and expression
	const (
		min int = 1 + 1
		max     // int = 1 + 1
	)

	fmt.Println(min, max)

	// print the types of min and max
	fmt.Printf("%T %T\n", min, max)
}

func main7() {
	// min and pi are typeless constants
	const min = 1 + 1
	const pi = 3.14 * min

	fmt.Println(min, pi)
}

//------------------- typed vs typeless constant

func main8() {
	const min int = 42

	var i int
	i = min // OK

	fmt.Println(i)
}

func main9() {
	const min int = 42

	var f float64

	// ERROR: Type Mismatch
	// f = min // NOT OK

	fmt.Println(f)
}

func main10() {
	const min = 42

	var f float64

	f = min // OK when min is typeless

	fmt.Println(f)
}

func main11() {
	const min = 42

	// I've removed int from the below declaration
	// since, min's default type is int (you'll learn)
	var i = min

	var f float64 = min
	var b byte = min
	var j int32 = min
	var r rune = min

	// behind the scenes:
	// below statement equals to:
	//
	// var b byte = min
	b = byte(min)

	fmt.Println(i, f, b, j, r)
}

func main12() {
	const min int32 = 1

	max := 5 + min
	// above line equals to this:
	// max := int32(5) + min

	fmt.Printf("Type of max: %T\n", max)
}

func main13() {
	const min = 1

	max := 5 + min
	// above line equals to this:
	// max := int(5) + int(min)

	fmt.Printf("Type of max: %T\n", max)
}

func main14() {
	i := 42
	f := 3.14
	b := true
	s := "Hello"
	r := 'A'

	fmt.Printf("i is %T\n", i)
	fmt.Printf("f is %T\n", f)
	fmt.Printf("b is %T\n", b)
	fmt.Printf("s is %T\n", s)

	// int32 = rune (type alias)
	fmt.Printf("r is %T\n", r)
}

func main15() {
	i := 42 + 3.14 // OK: 42 and 3.14 are typeless

	// above line equals to this:
	// i := float64(42 + 3.14)

	fmt.Printf("i is %T\n", i)
}

func main16() {
	// NOT OK
	// "Hello" string literal cannot be converted
	//   into a bool value
	// f := true + "Hello"
}

func main17() {
	type Duration int64

	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)

	// The types are main.Duration instead of just Duration
	// It's because, Duration type now belongs to main func
	fmt.Printf("Nanosecond is %T\n", Nanosecond)
	fmt.Printf("Microsecond is %T\n", Microsecond)
	fmt.Printf("Millisecond is %T\n", Millisecond)
	fmt.Printf("Second is %T\n", Second)
	fmt.Printf("Minute is %T\n", Minute)
	fmt.Printf("Hour is %T\n", Hour)

	// Print the types of time.Duration constants directly
	// This time, types will be time.Duration
	// It's because, they're inside the time package
	fmt.Printf("Nanosecond is %T\n", time.Nanosecond)
	fmt.Printf("Microsecond is %T\n", time.Microsecond)
	fmt.Printf("Millisecond is %T\n", time.Millisecond)
	fmt.Printf("Second is %T\n", time.Second)
	fmt.Printf("Minute is %T\n", time.Minute)
	fmt.Printf("Hour is %T\n", time.Hour)
}

func main18() {
	// #1
	t := time.Second * 10

	fmt.Println(t)

	// #2
	i := 10

	// NOT OK
	// time.Duration and int are different types
	// t = time.Second * i

	// OK: i is int, Duration is int64
	//     So, i is convertable to int64
	t = time.Second * time.Duration(i)

	fmt.Println(t)

	// #3
	const c = 10

	// OK: Because, c is typeless
	t = time.Second * c

	fmt.Println(t)
}

// 1 3.14 2.0.1 true 65
// 1 3.14 2.0.1 true 65
// ------------
// 2 6.28 2.0.1-beta false 66
// ------------
// 1 1000
// int int
// ------------
// 1 1
// int int
// ------------
// 2 2
// int int
// ------------
// 2 6.28
// ------------
// 42
// ------------
// 0
// ------------
// 42
// ------------
// 42 42 42 42 42
// ------------
// Type of max: int32
// ------------
// Type of max: int
// ------------
// i is int
// f is float64
// b is bool
// s is string
// r is int32
// ------------
// i is float64
// ------------
// ------------
// Nanosecond is main.Duration
// Microsecond is main.Duration
// Millisecond is main.Duration
// Second is main.Duration
// Minute is main.Duration
// Hour is main.Duration
// Nanosecond is time.Duration
// Microsecond is time.Duration
// Millisecond is time.Duration
// Second is time.Duration
// Minute is time.Duration
// Hour is time.Duration
// ------------
// 10s
// 10s
// 10s
