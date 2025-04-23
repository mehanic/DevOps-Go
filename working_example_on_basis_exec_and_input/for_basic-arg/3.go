package main

import "fmt"

func main() {
	n := 5
	a := []int{}
	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)
		a = append(a, k)
	}
	fmt.Println(a)
}
