package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c"}
	for i := range arr {
		fmt.Println(i)
	}
}

// 0
// 1
// 2
