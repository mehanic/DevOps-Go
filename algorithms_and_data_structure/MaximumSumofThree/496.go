package main

import "fmt"

func main() {
	n := 4
	arr := []int{1, 2, 3, 4}
	maxi := 0
	for i := 0; i < n; i++ {
		sumi := 0
		if i == 0 {
			sumi = arr[i] + arr[i+1] + arr[n-1]
		} else if i == n-1 {
			sumi = arr[i] + arr[0] + arr[n-2]
		} else {
			sumi = arr[i] + arr[i-1] + arr[i+1]
		}
		if sumi > maxi {
			maxi = sumi
		}
	}
	fmt.Println(maxi)
}
