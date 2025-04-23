package main

import (
	"fmt"
	"sort"
)

// 1. Recursive Binary Search
// Recursive Binary Search
func binarySearch(l, r int, nums []int, target int) int {
	if l > r {
		return -1 // Base case: target not found
	}
	m := l + (r-l)/2 // Find the middle index

	if nums[m] == target {
		return m // Target found at index m
	}
	if nums[m] < target {
		return binarySearch(m+1, r, nums, target) // Search right half
	}
	return binarySearch(l, m-1, nums, target) // Search left half
}

// Wrapper function for binary search
func search(nums []int, target int) int {
	return binarySearch(0, len(nums)-1, nums, target)
}

// Main function to test the binary search
func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15} // Sorted array
	target1 := 7
	target2 := 2

	fmt.Printf("Index of %d: %d\n", target1, search(arr, target1)) // Expected: 3
	fmt.Printf("Index of %d: %d\n", target2, search(arr, target2)) // Expected: -1 (not found)
	fmt.Println("-------")
	main1()
	fmt.Println("-------")
	main2()
	fmt.Println("-------")
	main3()
	fmt.Println("-------")
	main4()
}

// 2. Iterative Binary Search
func search1(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2

		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}

func main1() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15} // Sorted array
	target1 := 7
	target2 := 2

	fmt.Printf("Index of %d: %d\n", target1, search1(arr, target1)) // Expected: 3
	fmt.Printf("Index of %d: %d\n", target2, search1(arr, target2)) // Expected: -1 (not found)
}

// 3. Upper Bound
func search2(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		m := l + (r-l)/2
		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	if l > 0 && nums[l-1] == target {
		return l - 1
	}
	return -1
}

func main2() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15} // Sorted array
	target1 := 7
	target2 := 2

	fmt.Printf("Index of %d: %d\n", target1, search2(arr, target1)) // Expected: 3
	fmt.Printf("Index of %d: %d\n", target2, search2(arr, target2)) // Expected: -1 (not found)
}

////

// 4. Lower Bound
func search3(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		m := l + (r-l)/2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if l < len(nums) && nums[l] == target {
		return l
	}
	return -1
}

func main3() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15} // Sorted array
	target1 := 7
	target2 := 2

	fmt.Printf("Index of %d: %d\n", target1, search3(arr, target1)) // Expected: 3
	fmt.Printf("Index of %d: %d\n", target2, search3(arr, target2)) // Expected: -1 (not found)
}

//

func search4(nums []int, target int) int {
	index := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	if index < len(nums) && nums[index] == target {
		return index
	}
	return -1
}

func main4() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15} // Sorted array
	target1 := 7
	target2 := 2

	fmt.Printf("Index of %d: %d\n", target1, search4(arr, target1)) // Expected: 3
	fmt.Printf("Index of %d: %d\n", target2, search4(arr, target2)) // Expected: -1 (not found)
}
