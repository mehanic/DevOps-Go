package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	n := len(arr)
	sumi := 0
	maxi := 0

	// Loop through each element of the array
	for i := 0; i < n; i++ {
		// If we are at the last element, wrap around and sum with the first two elements
		if i == n-1 {
			sumi = arr[i] + arr[0] + arr[1]
			// If we are at the second-last element, wrap around and sum with the first element
		} else if i == n-2 {
			sumi = arr[i] + arr[i+1] + arr[0]
			// Normal case: sum the current element with the next two
		} else {
			sumi = arr[i] + arr[i+1] + arr[i+2]
		}
		// Update maxi if we find a larger sum
		if sumi > maxi {
			maxi = sumi
		}
	}
	// Print the maximum sum of 3 consecutive elements
	fmt.Println(maxi)
}

// 6
