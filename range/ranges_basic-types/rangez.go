package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 8}
	// range
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Println("-----------------")
	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Println("--------------------")
	// map
	m := map[string]int{"foo": 100, "baa": 200}
	// map
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// foo 100
// baa 200
