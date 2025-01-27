package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	Examples()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", Fib(i))
	}
	fmt.Println()
}

//---

// Examples demonstrates some of the functions
// in the math package
func Examples() {
	//sqrt Examples
	i := 25

	// i is an int, so convert
	result := math.Sqrt(float64(i))

	// sqrt of 25 == 5
	fmt.Println(result)

	// ceil rounds up
	result = math.Ceil(9.5)
	fmt.Println(result)

	// floor rounds down
	result = math.Floor(9.5)
	fmt.Println(result)

	// math also stores some consts:
	fmt.Println("Pi:", math.Pi, "E:", math.E)
}

//---

// global to memoize fib
var memoize map[int]*big.Int

func init() {
	// initialize the map
	memoize = make(map[int]*big.Int)
}

// Fib prints the nth digit of the fibonacci sequence
// it will return 1 for anything < 0 as well...
// it's calculated recursively and use big.Int since
// int64 will quickly overflow
func Fib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}

	// base case
	if n < 2 {
		memoize[n] = big.NewInt(1)
	}

	// check if we stored it before
	// if so return with no calculation
	if val, ok := memoize[n]; ok {
		return val
	}

	// initialize map then add previous 2 fib values
	memoize[n] = big.NewInt(0)
	memoize[n].Add(memoize[n], Fib(n-1))
	memoize[n].Add(memoize[n], Fib(n-2))

	// return result
	return memoize[n]
}

// 5
// 10
// 9
// Pi: 3.141592653589793 E: 2.718281828459045
// 1 1 2 3 5 8 13 21 34 55
