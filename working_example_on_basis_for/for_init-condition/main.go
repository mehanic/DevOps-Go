package main

import "fmt"

func main() {
	// we initialize i to 0, and use postincrement (no preincrement)
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}
