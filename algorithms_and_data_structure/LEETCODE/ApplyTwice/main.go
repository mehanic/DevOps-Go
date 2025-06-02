package main

import "fmt"

func applyTwice(fn func(int) int) func(int) int {
	return func(x int) int {
		return fn(fn(x))
	}
}

func main() {
	double := func(n int) int { return n * 2 }
	doubleTwice := applyTwice(double)
	fmt.Println(doubleTwice(3)) // double(double(3)) => 3 * 2 * 2 = 12

}
