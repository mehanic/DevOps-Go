package main

import "fmt"

func main() {
	arr := []int{100, 12, 233, 2324}
	fmt.Println("loop")
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
	fmt.Println("range")
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// loop
// 0 100
// 1 12
// 2 233
// 3 2324
// range
// 0 100
// 1 12
// 2 233
// 3 2324
