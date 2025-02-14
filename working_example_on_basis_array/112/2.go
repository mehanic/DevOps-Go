package main

import "fmt"

func addElement(arr []int) {
	arr[0] = 123
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)
	addElement(arr)
	fmt.Println(arr)
}

// ╰─λ go run 2.go                                                                                         0 (0.001s) < 21:27:01
// [1 2 3 4]
// [123 2 3 4]
