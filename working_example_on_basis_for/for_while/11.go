package main

import "fmt"

func main() {
	i := 0
	//while
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

// ╰─λ go run 11.go                                           0 (0.001s) < 17:53:58
// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
