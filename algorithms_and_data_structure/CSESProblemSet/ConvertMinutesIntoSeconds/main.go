package main

import (
	"fmt"
)

func convert(minutes int) int {
	return minutes * 60
}

func main() {
	fmt.Println(convert(5)) // ➞ 300
	fmt.Println(convert(3)) // ➞ 180
	fmt.Println(convert(2)) // ➞ 120
}
