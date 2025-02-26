package main

import "fmt"

func sortArr(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr
}

func main() {
	arr := []int{100, 223, 1, 2, 4, 51}
	arr = sortArr(arr)
	fmt.Println(arr)
}
