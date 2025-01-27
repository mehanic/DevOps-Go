package main

import (
	"fmt"
	"math"
)

// This program uses magic values

func main() {
	cm := 100
	m := cm / 100 // 100 is a magic value
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / 100 // 100 is a magic value
	fmt.Printf("%dcm is %dm\n", cm, m)

	main1()
	fmt.Println("------------------")
	//main2()
	//fmt.Println("------------------")
	main3()
	fmt.Println("------------------")
	main4()
	fmt.Println("------------------")
	main5()
	fmt.Println("------------------")
	main6()
}

// This program uses a named constant
// instead of a magic value

func main1() {
	const meters int = 100

	cm := 100
	m := cm / meters // using a named constant
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / meters // using a named constant
	fmt.Printf("%dcm is %dm\n", cm, m)
}

func main2() {
	// Go can't catch the same error at runtime
	// When you run this, there will be an error:
	//
	// panic: runtime error: integer divide by zero
	n, m := 1, 0
	fmt.Println(n / m)

	// Go will detect the division by zero error
	// at compile-time
	//
	// const n int = 1
	// const m int = 0
	// fmt.Println(n / m)
}

func main3() {
	const max int = 100

	// ERROR: cannot assign to max
	//        constants are immutable (cannot change)
	// max = 100
}

func main4() {
	// math.Pow calculates the power of the given number
	fmt.Println(math.Pow10(2)) // 100
	fmt.Println(math.Pow10(3)) // 1000
	fmt.Println(math.Pow10(4)) // 10000

	// ERROR: math.Pow is not a constant
	//        constants cannot use runtime constructs

	// const max int = math.Pow10(2)
}

func main5() {
	n := 2

	// ERROR: n is not a constant
	// const max int = n

	_ = n
}

func main6() {
	// When argument to len is a constant, len can be used
	// while initializing a constant
	//
	// Here, "Hello" is a constant.

	const max int = len("Hello")

	fmt.Println(max)
}

// 100cm is 1m
// 200cm is 2m
// 100cm is 1m
// 200cm is 2m
// ------------------
// ------------------
// 100
// 1000
// 10000
// ------------------
// ------------------
// 5
