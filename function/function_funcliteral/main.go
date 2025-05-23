package main

import "fmt"

func main() {
	func(n int) {
		fmt.Println(n)
	}(5)

	printslice := func(strs []string) {
		for _, str := range strs {
			fmt.Println(str)
		}
	}
	printslice([]string{"hello", "js", "world"})
}

// 5
// hello
// js
// world
