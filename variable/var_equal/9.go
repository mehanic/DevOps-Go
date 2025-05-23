package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	if k%2 == 0 {
		fmt.Println("четная")
	} else {
		fmt.Println("не четная")
	}
}

// 4
// четная

// 7
// не четная
