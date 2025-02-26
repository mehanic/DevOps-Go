package main

import "fmt"

func printMaxi(arr []int) {
	maxi := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxi {
			maxi = arr[i]
		}
	}
	fmt.Println(maxi)
}

func printSum(arr []int) {
	sumi := 0
	for i := 0; i < len(arr); i++ {
		sumi += arr[i]
	}
	fmt.Println(sumi)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{3, 4, 123, 1231, 5421}
	printMaxi(arr)
	printMaxi(arr2)
	printSum(arr)
}
