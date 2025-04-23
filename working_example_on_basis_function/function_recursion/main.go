package main

import "fmt"

func main() {
	r, s := factorial(5)
	fmt.Println(s, "=", r)
}

func factorial(n int) (int, string) {

	if n == 0 {
		return 1, "1"
	}
	r, s := factorial(n - 1)
	return n * r, fmt.Sprintf("%d * %s", n, s)
}
