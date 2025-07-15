package main

import "fmt"

func main() {
	var n int
	fmt.Println("input number of quality:")
	fmt.Scanf("%d", &n)
	arr := [100]int{}
	fmt.Printf("enter %d numbers:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	maxi := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > maxi {
			maxi = arr[i]
		}
	}
	fmt.Println("Maxi number:", maxi)
}
