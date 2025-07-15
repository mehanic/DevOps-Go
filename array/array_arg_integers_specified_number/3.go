package main

import (
	"fmt"
)

func main() {
	var n int = 5
	arr := [10]int{}

	values := [5]int{1, 2, 7, 5, 8}
	for i := 0; i < n; i++ {
		arr[i] = values[i]
	}

	fmt.Println(arr)
	k := len(arr)
	fmt.Println(k)
	sumi := 0
	for i := 0; i < n; i++ {
		sumi += arr[i]
	}
	fmt.Println(sumi)
}
