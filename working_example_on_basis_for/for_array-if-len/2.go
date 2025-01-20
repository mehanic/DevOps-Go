package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, -1, 3, 2}
	n := len(arr)
	//code
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	//
	fmt.Println(arr)
}
