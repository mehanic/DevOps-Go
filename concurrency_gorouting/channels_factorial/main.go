package main

import (
	"fmt"
	"strings"
)

func factorial(n int, ch chan int) {
	f := 1

	for i := 2; i <= n; i++ {
		f *= i
	}

	ch <- f
}

func main() {
	ch := make(chan int)
	defer close(ch)

	go factorial(5, ch)

	f := <-ch
	fmt.Println(f)

	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f = <-ch
		fmt.Println(f)
	}

	fmt.Println(strings.Repeat("#", 20))

	for i := 1; i <= 15; i++ {
		go func(n int, c chan int) {
			f := 1

			for i := 2; i <= n; i++ {
				f *= i
			}

			c <- f
		}(i, ch)

		fmt.Printf("Factorial of %d is %d\n", i, <-ch)
	}

}

// ─λ go run main.go                                                                        0 (0.001s) < 13:35:14
// 120
// 1
// 2
// 6
// 24
// 120
// 720
// 5040
// 40320
// 362880
// 3628800
// 39916800
// 479001600
// 6227020800
// 87178291200
// 1307674368000
// 20922789888000
// 355687428096000
// 6402373705728000
// 121645100408832000
// 2432902008176640000
// ####################
// Factorial of 1 is 1
// Factorial of 2 is 2
// Factorial of 3 is 6
// Factorial of 4 is 24
// Factorial of 5 is 120
// Factorial of 6 is 720
// Factorial of 7 is 5040
// Factorial of 8 is 40320
// Factorial of 9 is 362880
// Factorial of 10 is 3628800
// Factorial of 11 is 39916800
// Factorial of 12 is 479001600
// Factorial of 13 is 6227020800
// Factorial of 14 is 87178291200
// Factorial of 15 is 1307674368000
