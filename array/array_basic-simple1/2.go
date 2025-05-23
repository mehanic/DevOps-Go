package main

import "fmt"

func main() {
	var numbers [10]int
	fmt.Println(numbers)
	elements := [5]int{1, 2, 3, 4, 5}
	fmt.Println(elements)
	array := [...]int{1, 2, 4, 5, 6, 7, 8, 9, 2, 3, 4}
	fmt.Println(array)
}
