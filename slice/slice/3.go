package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 10}
	fmt.Println(len(sl))
	fmt.Println(cap(sl))
	lk := make([]int, 10)
	fmt.Println(len(lk))
	fmt.Println(cap(lk))
	lk = append(lk, 3)
}

// 7
// 7
// 10
// 10
