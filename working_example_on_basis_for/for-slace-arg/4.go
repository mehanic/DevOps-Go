package main

import "fmt"

func main() {
	n := 5
	a := []int{}
	fmt.Println(len(a))
	for i := 0; i < n; i++ {
		k := 5
		a = append(a, k)
	}
	fmt.Println(len(a))
}
