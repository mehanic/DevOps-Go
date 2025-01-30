package main

import "fmt"

func main() {
	// when an integer and a float value used together
	// in an expression, the result always becomes
	// a float value
	fmt.Println(8 * -4.0) // -32.0 not -32

	// two integer values result in an integer value
	fmt.Println(-4 / 2)

	// remainder operator
	// it can only used with integers
	fmt.Println(5 % 2)
	// fmt.Println(5.0 % 2) // wrong

	// addition operators
	fmt.Println(1 + 2.5)
	fmt.Println(2 - 3)

	// negation operator
	fmt.Println(-(-2))
	fmt.Println(-2) // this also works

	//------------------
	var (
		myAge   = 30
		yourAge = 35
		average float64
	)

	average = float64(myAge+yourAge) / 2

	fmt.Println(average)

	//----------------

	ratio := 1.0 / 10.0

	for range [...]int{10: 0} {
		ratio += 1.0 / 10.0
	}

	fmt.Printf("%.60f", ratio)

	//----
}

func main1() {

	ratio := 3 / 2

	fmt.Printf("%d", ratio)
}

func main2() {
	// When you use a float value with an integer
	// in a calculation,
	// the result always becomes a float.

	ratio := 3.0 / 2

	// OR:
	// ratio = 3 / 2.0

	// OR - if 3 is inside an int variable:
	// n := 3
	// ratio = float64(n) / 2

	fmt.Printf("%f", ratio)
}

func main3() {
	fmt.Println("sum :", 3+2)   // sum - int
	fmt.Println("sum :", 2+3.5) // sum - float64
	fmt.Println("dif :", 3-1)   // difference - int
	fmt.Println("dif :", 3-0.5) // difference - float64
	fmt.Println("prod:", 4*5)   // product - int
	fmt.Println("prod:", 5*2.5) // product - float64
	fmt.Println("quot:", 8/2)   // quotient - int
	fmt.Println("quot:", 8/1.5) // quotient - float64

	// remainder is only for integers
	fmt.Println("rem :", 8%3)
	// fmt.Println("rem:", 8.0%3) // error

	// you can do this
	// since the fractional part of a float is zero
	f := 8.0
	fmt.Println("rem :", int(f)%3)
}

func main4() {
	// what's the value of the ratio?
	// 3 / 2 = 1.5?
	var ratio float64 = 3 / 2
	fmt.Println(ratio)

	// explain
	// above expression equals to this:
	ratio = float64(int(3) / int(2))
	fmt.Println(ratio)

	// how to fix it?
	//
	// remember, when one of the values is a float value
	// the result becomes a float
	ratio = float64(3) / 2
	fmt.Println(ratio)

	// or
	ratio = 3.0 / 2
	fmt.Println(ratio)
}

func main5() {
	fmt.Println(
		2+2*4/2,
		2+((2*4)/2), // same as above
	)

	fmt.Println(
		1+4-2,
		(1+4)-2, // same as above
	)

	fmt.Println(
		(2+2)*4/2,
		(2+2)*(4/2), // same as above
	)
}

func main6() {
	n, m := 1, 5

	fmt.Println(2 + 1*m/n)
	fmt.Println(2 + ((1 * m) / n)) // same as above

	// let's change the precedence using parentheses
	fmt.Println(((2 + 1) * m) / n)
}

func main7() {
	celsius := 35.

	// Wrong formula  :  9*celsius + 160  / 5
	// Correct formula: (9*celsius + 160) / 5
	fahrenheit := (9*celsius + 160) / 5

	fmt.Printf("%g ºC is %g ºF\n", celsius, fahrenheit)
}

func main8() {
	var n int

	// ALTERNATIVES:
	// n = n + 1
	// n += 1

	// BETTER:
	n++

	fmt.Println(n)
}


// -32
// -2
// 1
// 3.5
// -1
// 2
// -2
// 32.5
// 1.199999999999999955591079014993738383054733276367187500000000%