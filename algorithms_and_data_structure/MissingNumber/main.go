// Missing Number
// Given an array nums containing n integers in the range [0, n] without any duplicates, return the single number in the range that is missing from nums.

// Follow-up: Could you implement a solution using only O(1) extra space complexity and O(n) runtime complexity?

// Example 1:

// Input: nums = [1,2,3]

// Output: 0
// Explanation: Since there are 3 numbers, the range is [0,3]. The missing number is 0 since it does not appear in nums.

// Example 2:

// Input: nums = [0,2]

// Output: 1
// Constraints:

// 1 <= nums.length <= 1000

package main

import "fmt"

func missingNumber(nums []int) int {
    n := len(nums)
    
    // Calculate the sum of the first n natural numbers (0 to n)
    totalSum := (n * (n + 1)) / 2
    
    // Calculate the sum of numbers in the array
    arraySum := 0
    for _, num := range nums {
        arraySum += num
    }
    
    // The missing number is the difference between the total sum and the array sum
    return totalSum - arraySum
}

func main() {
    // Example 1
    nums1 := []int{1, 2, 3}
    fmt.Println("Missing number:", missingNumber(nums1)) // Output: 0

    // Example 2
    nums2 := []int{0, 2}
    fmt.Println("Missing number:", missingNumber(nums2)) // Output: 1

    // Example 3
    nums3 := []int{0, 1}
    fmt.Println("Missing number:", missingNumber(nums3)) // Output: 2

    // Example 4
    nums4 := []int{0, 1, 3, 4}
    fmt.Println("Missing number:", missingNumber(nums4)) // Output: 2

    // Example 5
    nums5 := []int{0, 1, 2, 3, 5}
    fmt.Println("Missing number:", missingNumber(nums5)) // Output: 4
}


