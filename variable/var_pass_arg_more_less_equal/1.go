package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a > b {
		fmt.Println("a>b")
	} else if a == b {
		fmt.Println("a==b")
	} else {
		fmt.Println("a<b")
	}
}

// 4
// 7
// a<b
