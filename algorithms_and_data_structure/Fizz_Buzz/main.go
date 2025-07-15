package main

import (
	"fmt"
)

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			fmt.Println(i)
		} else {
			fmt.Println(s)
		}
	}
}

func main() {
	FizzBuzz(50)
}
