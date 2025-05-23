package main

import "fmt"

func main() {
	var people1 = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}

	fmt.Println(people1["Alice"]) // 8
	fmt.Println(people1["Bob"])   // 2

	people1["Bob"] = 32
	fmt.Println(people1["Bob"]) // 32
}

// 8
// 2
// 32
