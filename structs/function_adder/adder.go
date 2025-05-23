package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	// a := adder() is trivial and also works.
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n",
			i, s)
	}
}


// 0 + 1 + ... + 0 = 0
// 0 + 1 + ... + 1 = 1
// 0 + 1 + ... + 2 = 3
// 0 + 1 + ... + 3 = 6
// 0 + 1 + ... + 4 = 10
// 0 + 1 + ... + 5 = 15
// 0 + 1 + ... + 6 = 21
// 0 + 1 + ... + 7 = 28
// 0 + 1 + ... + 8 = 36
// 0 + 1 + ... + 9 = 45
