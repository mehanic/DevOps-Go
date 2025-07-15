package main

import (
	"fmt"
)

func shellSort(arr []int) {
	n := len(arr)
	gap := n / 2

	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
		gap /= 2
	}
}

func main() {
	list := []int{16, 2, 31, 45, 32, 13, 120, 28}
	shellSort(list)
	fmt.Println(list)
}
