package main

import "fmt"

func main() {

	visit([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	})
}

func visit(xi []int, callback func(int)) {
	for _, v := range xi {
		callback(v)
	}
}

// callback: passing a func as an argument
