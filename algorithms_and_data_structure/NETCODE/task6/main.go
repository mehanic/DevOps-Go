// Contains Duplicate
// Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

// Example 1:

// Input: nums = [1, 2, 3, 3]

// Output: true

// Example 2:

// Input: nums = [1, 2, 3, 4]

// Output: false

// Recommended Time & Space Complexity
// You should aim for a solution with O(n) time and O(n) space, where n is the size of the input array.

// Hint 1
// A brute force solution would be to check every element against every other element in the array. This would be an O(n^2) solution. Can you think of a better way?

// Hint 2
// Is there a way to check if an element is a duplicate without comparing it to every other element? Maybe there's a data structure that is useful here.

// Hint 3
// We can use a hash data structure like a hash set or hash map to store elements we've already seen. This will allow us to check if an element is a duplicate in constant time.

package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool) // Create a hash map to store the elements we've seen

	for _, num := range nums {
		if seen[num] {
			return true // If the number is already in the map, return true (duplicate found)
		}
		seen[num] = true // Add the number to the map
	}

	return false // No duplicates found
}

func main() {
	// Example 1: Duplicate found
	nums1 := []int{1, 2, 3, 3}
	fmt.Println(containsDuplicate(nums1)) // Output: true

	// Example 2: No duplicates
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums2)) // Output: false

	// Example 3: Duplicate found
	nums3 := []int{5, 6, 7, 8, 9, 6}
	fmt.Println(containsDuplicate(nums3)) // Output: true

	// Example 4: No duplicates
	nums4 := []int{10, 20, 30, 40, 50}
	fmt.Println(containsDuplicate(nums4)) // Output: false

	fmt.Println("-----")
	main1()
	fmt.Println("-----")
	main2()

}

// other variant
func containsDuplicate1(nums []int) bool {
	seen := make(map[int]int) // Create a map to store counts of the elements we've seen

	for _, num := range nums {
		if seen[num] > 0 { // If the number has been seen before (count > 0), return true
			return true
		}
		seen[num] = 1 // Set the count of the number to 1 (indicating it's been seen once)
	}

	return false // No duplicates found
}

func main1() {
	nums := []int{5, 5, 5, 5, 5}
	fmt.Println(containsDuplicate1(nums)) // Output: true

	nums1 := []int{1, 2, 3, 4, 5}
	fmt.Println(containsDuplicate(nums1)) // Output: false

	nums2 := []int{1}
	fmt.Println(containsDuplicate(nums2)) // Output: false
}

func containsDuplicate2(nums []int) bool {
	seen := make(map[int]int) // Map to store counts of numbers

	for _, num := range nums {
		if seen[num] > 0 { // If the count is greater than 0, we've already seen this number
			return true
		}
		seen[num] = seen[num] + 1 // Increment the count of the number
	}

	return false
}

func main2() {
	nums := []int{}
	fmt.Println(containsDuplicate2(nums)) // Output: false

	nums1 := []int{5, 5, 5, 5, 5}
	fmt.Println(containsDuplicate2(nums1)) // Output: true

	nums2 := []int{1, 2, 3, 4, 5}
	fmt.Println(containsDuplicate2(nums2)) // Output: false

	nums3 := []int{1, 2, 3, 4, 2}
	fmt.Println(containsDuplicate2(nums3)) // Output: true
}
